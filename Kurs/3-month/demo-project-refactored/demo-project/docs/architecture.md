# Архитектура

Как устроен проект и почему так. Как добавить фичу — в [adding-feature.md](adding-feature.md), команды — в [commands.txt](commands.txt), API — в [openapi.yaml](openapi.yaml).

## 1. Слои

```
Клиент
  → middleware: Logging → Recover → Auth
  → handler/v1     разобрать HTTP → вызвать сервис → ответить
  → service        бизнес-правила и права
  → repository     SQL, Redis, S3
      Postgres: roles, users, posts, sessions, files
      Redis:    session:<id>:access, session:<id>:refresh (хеши токенов)
      MinIO:    сами файлы
```

| Слой | Делает | Не делает |
|---|---|---|
| `handler/v1` | читает JSON, путь, заголовки; собирает вход; отвечает JSON | бизнес-проверки, SQL |
| `service` | права, `in.Validate()`, бизнес-сценарий | HTTP, JSON, SQL, логи |
| `repository` | один запрос или одна транзакция; ошибки БД → ошибки `models` | бизнес-решения |
| `models` | структуры, ошибки, простые правила (`CurrentUser.CanCreate`, `Input.Validate`) | ввод-вывод |
| `cron` | расписание фоновой очистки и логи | бизнес-логику |

```
cmd/api/main.go     app.Run()
internal/
  app/              сборка: config → Postgres/Redis/S3 → repositories → services → router → server
  config/           настройки из env
  handler/          router.go, v1/ (ручки и DTO), middleware/, response/
  service/          service.go (интерфейсы), auth/, user/, post/, file/, cleanup/
  repository/       repository.go (интерфейсы), user/, session/, token/, post/, file/, storage/
  cron/             расписание очистки (gocron v1)
  models/           структуры, входы с Validate, CurrentUser, роли, ошибки
pkg/                postgres/, redis/, s3/, logger/
migrations/         001_roles_users, 002_posts, 003_sessions, 004_files
```

Интерфейсы репозиториев — в `repository/repository.go`, сервисов — в `service/service.go`. Сборка — только в `internal/app/app.go`.

## 2. Как написан метод сервиса

1. `ctx`, затем `user models.CurrentUser` (если важно, кто действует), затем данные запроса: 1–2 значения — отдельными параметрами, больше — одной структурой `XxxInput`.
2. `in.Validate()` сначала приводит поля к виду (trim, lower-case), потом проверяет. Он лежит в `models` рядом со структурой.
3. Правила, для которых нужны БД, конфиг или права, — в сервисе.
4. Ошибки не глотаем: каждую возвращаем с контекстом (`fmt.Errorf("...: %w", err)`) или объединяем через `errors.Join`.

```go
func (s *Service) Update(ctx context.Context, user models.CurrentUser, in models.UpdatePostInput) error {
	if err := in.Validate(); err != nil {
		return err
	}

	post, err := s.getPostForEdit(ctx, user, in.ID)
	if err != nil {
		return fmt.Errorf("get post: %w", err)
	}

	post.Title = in.Title
	post.Content = in.Content
	if err := s.posts.Update(ctx, post); err != nil {
		return fmt.Errorf("update post: %w", err)
	}
	return nil
}
```

## 3. Бизнес-правила

**Пользователи**
- Регистрация — всегда с ролью `user`. Email приводится к нижнему регистру и должен быть уникальным.
- Пароль — от 8 символов и не больше 72 байт (ограничение bcrypt). Имя — от 1 до 100 символов.

**Сессии и токены**
- Login создаёт сессию (строка в `sessions`, id — UUID) и пару токенов. В Redis лежат SHA-256 хеши токенов, а не сами токены.
- На неизвестный email и неверный пароль ответ одинаковый — `401 invalid email or password`.
- Access-токен проходит, только если совпадает с текущим в Redis. После logout или refresh старый access не работает.
- Refresh-токен одноразовый. Refresh не создаёт новую сессию: та же сессия получает новую пару токенов и продлевается.
- Logout, «завершить сессию», «выйти на других устройствах» удаляют строки сессий и их ключи в Redis.

**Посты**
- Создавать посты могут все роли, кроме `auditor` и `support`.
- Менять и удалять пост может автор, `moderator` или `admin`.
- Заголовок — 1–200 символов, текст — 1–10 000. К посту — до 10 своих неприкреплённых файлов.
- Пост и привязка файлов создаются в одной транзакции: если хотя бы один файл чужой или уже прикреплён, пост не создаётся.
- Удаление поста удаляет и его файлы: записи — каскадом в БД, объекты — из S3.

**Файлы**
- До 10 МБ. Тип определяется по содержимому, а не по расширению: JPEG, PNG, WebP, GIF, PDF.
- Прикреплённый к посту файл видят все. Неприкреплённый — только владелец и `admin`, остальным `404`.
- Удалить можно только неприкреплённый файл (для прикреплённого — `409`). Удаляются и запись, и объект в S3.
- Файл, который не прикрепили к посту за `CLEANUP_UNATTACHED_TTL` (24 часа), удаляет фоновая очистка.

Права — методы `models.CurrentUser`: `CanCreate()`, `CanEditPost(authorID)`, `CanManageFile(ownerID)`.

## 4. Потоки

### Login

1. `LoginInput.Validate`: email в нижний регистр, `X-Device-Type` → `web` / `mobile` / `400`.
2. Найти пользователя по email и сверить пароль через bcrypt.
3. Создать сессию в Postgres.
4. `issueTokens`: подписать access (с `role_id`) и refresh, сохранить их хеши в Redis.
5. Если токены не сохранились — удалить только что созданную сессию.

### Refresh

1. Взять refresh из `Authorization: Bearer`, проверить подпись, срок и тип.
2. `UseRefresh`: команда Redis `GETDEL` забирает хеш refresh одной командой (прочитать + удалить). Из двух одновременных запросов с одним токеном пройдёт только первый.
3. Перечитать пользователя: в refresh нет роли, и она могла измениться.
4. Продлить сессию (`expires_at`) и выдать новую пару для той же сессии.

### Logout

Обычная защищённая ручка: `requireAuth` проверяет access, сервис удаляет строку сессии и её ключи в Redis.

### Пост с файлами

```go
tx, _ := db.Begin(ctx)
defer rollback(tx)

tx.QueryRow(`INSERT INTO posts ... RETURNING id`)
tag := tx.Exec(`UPDATE files SET post_id = $1
                 WHERE id = ANY($2) AND user_id = $3 AND post_id IS NULL`)
if tag.RowsAffected() != len(fileIDs) {
    return models.ErrFileNotAvailable // транзакция откатится
}
tx.Commit(ctx)
```

### Фоновая очистка

`internal/cron` раз в `CLEANUP_INTERVAL` (30 минут) запускает один проход с `SingletonMode`: следующий запуск не начнётся, пока не закончился текущий. Шаги (ошибка одного не мешает другому):

1. **Брошенные загрузки** — файлы без поста старше `CLEANUP_UNATTACHED_TTL`: удаляются запись и объект в S3. Если файл успели прикрепить между выборкой и удалением, он пропускается.
2. **Истёкшие сессии** — `DELETE FROM sessions WHERE expires_at < now()`. Их ключи в Redis к этому времени уже удалены по TTL.

## 5. Решения

| Решение | Почему |
|---|---|
| В Redis — хеши токенов, а не сами JWT | утечка Redis не даёт готовых токенов |
| Сессии — в Postgres, проверка токена — в Redis | список устройств хранится надёжно, а проверка на каждом запросе быстрая |
| id сессии и ключи файлов — случайные UUID | привычный формат; угадать чужой id нельзя |
| В access-токене `role_id`; id ролей совпадают с константами `models.Role*` | права проверяются без похода в БД |
| Refresh одноразовый, сессия та же | украденный refresh нельзя использовать дважды; не нужно переносить данные сессии |
| Выход из сессии — удаление строки | историю отзывов никто не читает; меньше колонок и состояний |
| Объекты S3 удаляются вместе с записью | не нужен фоновый поиск «сирот» |
| Bucket приватный, ссылки presigned | файлы недоступны без авторизации |
| Миграции — только итоговая схема (`CREATE`, без `ALTER`) | проект ещё не в production: схема поменялась — правим файл и `make infra-reset` |
| Без ORM, generic-репозиториев, библиотек валидации | обычный Go-код понятнее |

## 6. Компромиссы

- **S3 недоступен при удалении.** Запись уже удалена, клиент получит `500`, а объект останется в bucket.
- **Повторный старый refresh.** Если предъявить уже использованный refresh-токен, текущая пара этой сессии тоже сгорает: так выглядит кража токена, и безопаснее разлогинить.
- **Роль в access живёт до его истечения** (15 минут). Новая роль приходит после refresh.
- **Несколько копий API запустят очистку одновременно.** Удаления безопасны, но работа дублируется. Для production — распределённая блокировка.
- **Нет migration-tool.** Для production нужен goose или atlas.

## 7. За reverse proxy

Сейчас IP клиента берётся из соединения (`RemoteAddr`). Если поставить API за nginx или балансировщик, в сессиях будет IP proxy. Тогда:

1. Добавить настройку `HTTP_TRUST_PROXY=true` и передать её в `v1.Config`.
2. В `clientIP` (`internal/handler/v1/request.go`) при включённой настройке брать **последний** адрес из `X-Forwarded-For` — его дописал ваш proxy. Первые адреса пишет клиент, их легко подделать.
3. Включать только за доверенным proxy.
