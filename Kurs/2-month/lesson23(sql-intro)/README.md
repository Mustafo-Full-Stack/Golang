# Урок 23 — Введение в SQL и подключение к БД

Шпаргалка для лайв-кодинга. Полная теория, тайминг и практика — в `Urok_22_Vvedenie_v_SQL.md`.

## Главные аналогии (сказать в начале)
- База данных — **склад с идеальным кладовщиком**: «дай все ноутбуки дешевле 1000, отсортируй по цене» — приносит за миллисекунды. Файл — коробка на чердаке: чтобы что-то найти, надо перебрать всё вручную.
- `sql.DB` — это **пул соединений**, а не одно соединение. Как стойка с телефонами: берёшь свободный, поговорил — вернул на место.
- Плейсхолдер `$1` — **бланк с пустой графой**: текст запроса печатается заранее, а данные вписываются отдельно, и «превратиться в команду» они уже не могут.

## Подготовка (до урока)
1. Запустить PostgreSQL через Docker Compose (из корня репозитория, где лежит `docker-compose.yml`):
   ```bash
   docker compose up -d --wait
   docker compose ps          # STATUS: Up ... (healthy)
   ```
   Все команды compose (логи, остановка, сброс данных, ошибки) — в `compose-commands.txt` в корне репозитория.
2. Проверить подключение: `docker compose exec postgres psql -U postgres -d coursedb` (psql уже внутри контейнера, ставить на компьютер не нужно).
3. В папке урока уже лежат `go.mod` и `go.sum` с драйвером `lib/pq` (шаги 3–11) и `pgx` (шаг 13). Если хотите показать студентам установку вживую — временно удалите их и выполните:
   ```bash
   go mod init dbdemo
   go get github.com/lib/pq
   go get github.com/jackc/pgx/v5@v5.7.6   # нужен только для шага 13
   go mod tidy
   ```
   Версия `pgx` закреплена: самая свежая (v5.11) требует Go 1.25 и поднимет директиву `go` в `go.mod`, а в курсе используется Go 1.24.
4. Сброс базы к началу (если что-то сломали или повторяете урок): файл `reset.sql`
   ```bash
   docker compose exec -T postgres psql -U postgres -d coursedb < reset.sql
   # или, если psql установлен на компьютере:
   psql -h localhost -U postgres -d coursedb -f reset.sql
   ```

## Как запускать
Шаги 1–2 и 12 — SQL-файлы, команды копируем в `psql` по одной. Шаги 3–11 и 13 — Go-программы:
```bash
cd "lesson23(sql-intro)"
go run ./05-query-mnogo-strok
# или из папки шага: cd 05-query-mnogo-strok && go run main.go
```
Строка подключения в шагах 3–11 одна: `host=localhost port=5432 user=postgres password=secret dbname=coursedb sslmode=disable` — для урока так проще всего. Как делают в реальных проектах — в шаге 13 (настройки из переменных окружения, драйвер `pgx`; пример: `DB_PASSWORD=wrong go run ./13-podklyuchenie-best-practice`).

## Порядок шагов
| # | Папка | Что показываем |
|---|---|---|
| 1 | `01-sozdaem-tablicu` | `CREATE TABLE`, `INSERT`. Ограничения: UNIQUE и CHECK не пускают плохие данные. |
| 2 | `02-vybiraem-dannye` | `SELECT`: `WHERE`, `ORDER BY`, `LIMIT`, `COUNT/AVG`, `LIKE`, `BETWEEN`. |
| 3 | `03-podklyuchenie-k-baze` | `sql.Open` + `Ping`. Драйвер с `_`. Что будет при неверном пароле. |
| 4 | `04-nastroika-pula` | `SetMaxOpenConns` и другие, `db.Stats()`. `Open` ничего не открывает. |
| 5 | `05-query-mnogo-strok` | `db.Query` — шаблон из 5 пунктов: `Query → defer Close → Next → Scan → Err`. |
| 6 | `06-queryrow-odna-stroka` | `db.QueryRow` и `sql.ErrNoRows` («не нашлось» — это не поломка). |
| 7 | `07-parametry-i-inekcii` | `$1` против склейки строк. Живая SQL-инъекция `' OR '1'='1`. |
| 8 | `08-agregaty-i-null` | `COUNT/AVG/MIN/MAX`, `sql.NullFloat64`, `.Valid`. |
| 9 | `09-dinamicheskii-filtr` | Фильтр на указателях: запрос собирается из частей, значения — только через `args`. |
| 10 | `10-konsolnyi-prosmotrshchik` | Мини-проект: меню + база. Всё вместе. |
| 11 | `11-tipichnye-oshibki` | Ломаем нарочно: порядок в `Scan`, NULL в `float64`, забытый `rows.Close()`. |
| 12 | `12-update-delete` | `UPDATE` и `DELETE` в `psql`. Правило: сначала `SELECT` с тем же `WHERE`. Мягкое удаление. Забытый `WHERE` (закомментирован). **Показывать сразу после шага 2**, до Go-программ. |
| 13 | `13-podklyuchenie-best-practice` | Подключение как в настоящем проекте: структура `DBConfig`, настройки из переменных окружения (`DB_HOST`, `DB_PASSWORD`…), адрес собирает `net/url`, драйвер `pgx` + `pgxpool`, таймаут и повторные попытки, пароль не в логах. **Показывать после шага 4** или в конце как «а теперь как в проде». |

Что здесь **нарочно сломано** — папка `11-tipichnye-oshibki` (и «плохая» функция `findByNameUnsafe` в шаге 7). Не исправляйте.

## Частые ошибки новичков
| Ошибка | Симптом | Решение |
|---|---|---|
| Нет `_` перед драйвером | `unknown driver "postgres"` | `_ "github.com/lib/pq"` |
| Забыт `rows.Close()` | Утечка соединений, зависание | `defer rows.Close()` сразу после `Query` |
| Порядок в `Scan` | `Scan error ... converting ...` | Строго как в `SELECT` |
| NULL в обычный тип | `converting NULL to float64 is unsupported` | `sql.NullXxx` или указатель `*T` |
| Склейка SQL строками | SQL-инъекция | Плейсхолдеры `$1` |
| Нет `rows.Err()` | Тихая потеря данных | Проверять после цикла |
| `sql.Open` в каждой функции | Куча пулов | Один `*sql.DB` на приложение |
| `sslmode` не указан | `connection requires SSL` | `sslmode=disable` локально |
| Пароль в коде и в git | Пароль видят все, у кого есть репозиторий | Настройки из переменных окружения (шаг 13) |
| Нет таймаута на подключение | Программа «висит», если база не отвечает | `context.WithTimeout` + `Ping` (шаг 13) |
| `UPDATE`/`DELETE` без `WHERE` | Изменены или удалены **все** строки | Сначала `SELECT` с тем же `WHERE` (шаг 12) |
| Порт 5432 занят / `connection refused` | Контейнер не стартует или Go не подключается | `docker compose ps`, `lsof -i :5432`, см. `compose-commands.txt` |

## Вопросы группе в конце
1. Почему БД лучше файла для поиска?
2. Что делает `sql.Open` и почему нужен `Ping`?
3. Зачем `_` перед импортом драйвера?
4. Что такое SQL-инъекция и как Go от неё защищает?
5. Чем `Query` отличается от `QueryRow` и `Exec`?
