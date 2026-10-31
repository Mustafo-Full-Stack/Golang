# demo-project

Учебный API на Go: пользователи, роли, сессии на JWT, посты и файлы в S3/MinIO.

Стек: Go 1.27, `net/http`, PostgreSQL 17, Redis 7, MinIO, pgx, go-redis, JWT, gocron.

## Запуск

```bash
cp .env.example .env   # локальные настройки
make infra-up          # Postgres, Redis, MinIO; миграции применятся сами
make run               # API на http://localhost:8080
make test              # gofmt + go vet + go test
```

- MinIO: API — `http://localhost:9000`, консоль — `http://localhost:9001` (`minioadmin` / `minioadmin`).
- Логи: при `APP_ENV=local` — в файл `app.log`, иначе JSON в stdout.
- Миграции применяются только при **первом** создании базы. Поменяли SQL — `make infra-reset` (данные удалятся).

## Документация

| Файл | Что внутри |
|---|---|
| [docs/architecture.md](docs/architecture.md) | как устроен проект: слои, бизнес-правила, потоки, решения |
| [docs/adding-feature.md](docs/adding-feature.md) | как добавить новую фичу — пошагово |
| [docs/commands.txt](docs/commands.txt) | команды и curl-сценарии для проверки |
| [docs/openapi.yaml](docs/openapi.yaml) | контракт API |

## Ручки

| Ручка | Доступ | Что делает |
|---|---|---|
| `POST /api/v1/auth/register` | все | регистрация, роль всегда `user` |
| `POST /api/v1/auth/login` | все | создаёт сессию, возвращает пару токенов |
| `POST /api/v1/auth/refresh` | refresh-токен | новая пара токенов для той же сессии |
| `POST /api/v1/auth/logout` | access-токен | завершает текущую сессию |
| `GET /api/v1/users/me` | access-токен | профиль |
| `GET /api/v1/sessions` | access-токен | активные сессии пользователя |
| `DELETE /api/v1/sessions/{id}` | access-токен | завершить свою сессию |
| `DELETE /api/v1/sessions` | access-токен | выйти на всех других устройствах |
| `POST /api/v1/files` | пишущая роль | загрузить файл (multipart, поле `file`) |
| `GET /api/v1/files/{id}` | access-токен | файл и ссылка на скачивание |
| `DELETE /api/v1/files/{id}` | владелец, admin | удалить неприкреплённый файл |
| `GET /api/v1/posts` | access-токен | лента с файлами |
| `POST /api/v1/posts` | пишущая роль | создать пост, можно с `file_ids` |
| `GET /api/v1/posts/{id}` | access-токен | один пост |
| `PUT /api/v1/posts/{id}` | автор, moderator, admin | изменить пост |
| `DELETE /api/v1/posts/{id}` | автор, moderator, admin | удалить пост вместе с файлами |

Токены передаются в заголовке `Authorization: Bearer <token>`. Тип устройства при логине — в заголовке `X-Device-Type`: `web` (по умолчанию) или `mobile`.

## Роли

| Роль | Посты | Файлы |
|---|---|---|
| `user`, `employer`, `applicant` | читать все, менять свои | загружать, смотреть и удалять свои |
| `moderator` | читать все, менять любые | загружать, смотреть и удалять свои |
| `admin` | читать все, менять любые | любые |
| `auditor`, `support` | только читать | только смотреть прикреплённые к постам |

Служебные роли назначаются SQL-ом — см. `docs/commands.txt`.

## Конфиг

Все переменные — в `.env.example`. Главные:

| Переменная | По умолчанию | Зачем |
|---|---|---|
| `JWT_SECRET` | обязательна | не короче 32 байт |
| `JWT_ACCESS_TTL` / `JWT_REFRESH_TTL` | `15m` / `720h` | сроки жизни токенов |
| `S3_MAX_UPLOAD_BYTES` | `10485760` | максимальный размер файла (10 МБ) |
| `S3_PRESIGN_TTL` | `15m` | сколько живёт ссылка на файл |
| `CLEANUP_INTERVAL` | `30m` | как часто запускать фоновую очистку |
| `CLEANUP_UNATTACHED_TTL` | `24h` | через сколько удалить файл, который не прикрепили к посту |

## Ошибки

Тело ошибки всегда `{"error": "..."}`.

| Статус | Когда |
|---|---|
| 400 | неверный JSON или поля, недоступный `file_id` |
| 401 | нет токена, токен просрочен или отозван, неверный email или пароль |
| 403 | не хватает прав |
| 404 | не найдено (в том числе чужой неприкреплённый файл) |
| 409 | email занят, файл уже прикреплён к посту |
| 413 | файл больше лимита |
| 415 | тип файла не поддерживается |
| 500 | внутренняя ошибка, подробности — только в логе |
