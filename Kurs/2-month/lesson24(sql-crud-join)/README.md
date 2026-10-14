# Урок 24 — INSERT, UPDATE, DELETE и JOIN

Шпаргалка для лайв-кодинга. Полная теория, тайминг и практика — в `Urok_23_Zaprosy.md`.

## Главные аналогии (сказать в начале)
- `JOIN` — **склейка двух списков по общему признаку**. Есть список заказов (в нём только id клиента) и список клиентов. JOIN ставит рядом с заказом имя клиента, а не безликое число.
- `INNER JOIN` — «только те, у кого есть пара». `LEFT JOIN` — «все из левого списка, а если пары нет — пустое место (NULL)».
- Repository — **окошко в архиве**: вы говорите «дайте дело номер 5», а в какой шкаф за ним ходят — не ваша забота. Весь SQL живёт в одном месте.
- Забытый `WHERE` в `UPDATE`/`DELETE` — способ уничтожить продакшн за секунду. Всегда сначала проверяйте условие запросом `SELECT`.

## Подготовка (до урока)
1. PostgreSQL запущен (см. `lesson23(sql-intro)/README.md`), база `coursedb`, пароль `secret`.
2. Привести базу к началу урока — `reset.sql` (таблица `users` и 4 пользователя):
   ```bash
   docker compose exec -T postgres psql -U postgres -d coursedb < reset.sql
   # или, если psql установлен на компьютере:
   psql -h localhost -U postgres -d coursedb -f reset.sql
   ```
3. `go.mod` с драйвером `lib/pq` уже лежит в папке урока.

## Как запускать
```bash
cd "lesson24(sql-crud-join)"
go run ./04-update
# или из папки шага: cd 04-update && go run main.go
```
Шаги идут по порядку и опираются на данные друг друга. Если запускаете шаг повторно или что-то пошло не так — снова выполните `reset.sql`. Шаги 1 и 2 при повторном запуске дадут ошибку «email уже есть» — это нормально, и это же хороший повод перейти к шагу 3.

Файлы `.sql` в папках шагов 6, 7, 8 выполняем в `psql` **по ходу урока** (там создаются таблицы `tags` и `orders`):
```bash
psql -h localhost -U postgres -d coursedb -f 07-join/orders.sql
```

## Порядок шагов
| # | Папка | Что показываем |
|---|---|---|
| 1 | `01-insert-exec` | `INSERT` через `db.Exec`, `RowsAffected`. |
| 2 | `02-returning-id` | `RETURNING id` + `QueryRow().Scan()` — как получить ID в PostgreSQL. |
| 3 | `03-oshibki-bd` | `pq.Error` + `errors.As`: превращаем «unique violation» в понятное сообщение. |
| 4 | `04-update` | `UPDATE`: проверка `RowsAffected`, несколько полей, частичное обновление (`*string`, `*int`). |
| 5 | `05-delete` | `DELETE`, мягкое удаление (`is_active`) и восстановление. |
| 6 | `06-massovaya-vstavka` | 1000 вставок тремя способами: цикл → `Prepare` → один запрос. Сравниваем время. Сначала `tags.sql`. |
| 7 | `07-join` | Вторая таблица, внешний ключ, `INNER JOIN`. Сначала `orders.sql`. |
| 8 | `08-left-join` | `LEFT JOIN` + `GROUP BY` + `sql.NullFloat64`. Сначала в `psql`: `left-join.sql`. |
| 9 | `09-repository` | Паттерн Repository: полный CRUD, в `main` ни одного слова SQL. |
| 10 | `10-paginaciya` | `LIMIT/OFFSET`, общее количество, номер страницы. |
| 11 | `11-tipichnye-oshibki` | `LastInsertId`, `UPDATE` без `WHERE`, нарушение внешнего ключа, `Query` вместо `Exec`. Нужна таблица `orders` (шаг 7). |

Что здесь **нарочно сломано** — папка `11-tipichnye-oshibki`. Не исправляйте.

## Частые ошибки новичков
| Ошибка | Симптом | Решение |
|---|---|---|
| `LastInsertId` в PostgreSQL | `not supported by this driver` | `RETURNING id` |
| Нет `WHERE` в `UPDATE`/`DELETE` | Изменились все строки | Всегда проверять запрос через `SELECT` |
| `Query` для `INSERT` | Утечка соединений | Использовать `Exec` |
| Не проверен `RowsAffected` | «Обновилось» несуществующее | Проверять, что строк > 0 |
| Нет `defer stmt.Close()` | Утечка на сервере БД | Закрывать подготовленное выражение |
| SQL по всему проекту | Невозможно менять БД | Паттерн Repository |
| `INNER JOIN` вместо `LEFT` | Пропали записи без пары | Понимать разницу |

## Вопросы группе в конце
1. Как получить ID вставленной записи в PostgreSQL?
2. Почему `UPDATE` несуществующей строки не даёт ошибку?
3. Чем `INNER JOIN` отличается от `LEFT JOIN`?
4. Зачем нужен паттерн Repository?
5. Что делает внешний ключ и что такое `ON DELETE CASCADE`?
