-- Шаг 12. UPDATE и DELETE (выполняем в psql, команду за командой)
-- Показывать сразу после шага 2 (SELECT), ещё до Go-программ.
--
-- Подключение (база поднята через docker compose, см. compose-commands.txt):
--   docker compose exec postgres psql -U postgres -d coursedb
--
-- Начинаем с чистой базы (таблица users и 4 пользователя). Если что-то не так:
--   docker compose exec -T postgres psql -U postgres -d coursedb < reset.sql
--
-- Четыре главные команды SQL - их называют CRUD:
--   INSERT  - добавить строку   (Create)
--   SELECT  - прочитать         (Read)
--   UPDATE  - изменить          (Update)
--   DELETE  - удалить           (Delete)
--
-- ЗОЛОТОЕ ПРАВИЛО: перед UPDATE и DELETE сначала пишем SELECT с тем же WHERE.
-- Сначала смотрим, ЧТО заденем, и только потом меняем или удаляем.


-- =====================================================================
-- Что у нас сейчас
-- =====================================================================
SELECT * FROM users ORDER BY id;


-- =====================================================================
-- UPDATE - меняем данные
-- Шаблон: UPDATE таблица SET столбец = значение WHERE условие;
-- =====================================================================

-- 1) Меняем одно поле у одного человека.
--    Сначала смотрим, кого заденем:
SELECT * FROM users WHERE name = 'Тимур';
--    Теперь меняем:
UPDATE users SET age = 20 WHERE name = 'Тимур';
-- UPDATE 1   <- база сообщает, сколько строк изменилось
SELECT * FROM users WHERE name = 'Тимур';

-- 2) Несколько полей сразу - через запятую
UPDATE users SET name = 'Вера Ивановна', email = 'vera.i@mail.tj' WHERE id = 2;
SELECT * FROM users WHERE id = 2;

-- 3) Новое значение можно считать от старого (у Али день рождения)
UPDATE users SET age = age + 1 WHERE id = 1;
SELECT name, age FROM users WHERE id = 1;

-- 4) Изменить сразу много строк - всех, кто подходит под условие
SELECT * FROM users WHERE age > 25;
UPDATE users SET is_active = false WHERE age > 25;
-- UPDATE 3   <- изменилось 3 строки (Али 26, Вера 30, Нигина 27)
SELECT name, age, is_active FROM users ORDER BY id;

-- 5) Обновить строку, которой нет - это НЕ ошибка
UPDATE users SET age = 40 WHERE id = 9999;
-- UPDATE 0   <- запрос выполнился, но менять было нечего
--    Именно поэтому в Go после Exec проверяют RowsAffected (урок 24).

-- 6) База защищает данные и при UPDATE
UPDATE users SET age = 200 WHERE id = 1;
-- ERROR: new row for relation "users" violates check constraint "users_age_check"

UPDATE users SET email = 'ali@mail.tj' WHERE id = 2;
-- ERROR: duplicate key value violates unique constraint "users_email_key"


-- =====================================================================
-- DELETE - удаляем строки
-- Шаблон: DELETE FROM таблица WHERE условие;
-- =====================================================================

-- Добавим пару "лишних" записей, чтобы было что удалять (повторяем INSERT)
INSERT INTO users (name, email, age) VALUES
    ('Тест', 'test@mail.tj', 99),
    ('Временный', 'temp@mail.tj', 99);
SELECT * FROM users ORDER BY id;

-- 1) Удаляем одну строку. Сначала SELECT:
SELECT * FROM users WHERE email = 'test@mail.tj';
DELETE FROM users WHERE email = 'test@mail.tj';
-- DELETE 1
SELECT * FROM users ORDER BY id;

-- 2) Удаляем по условию (может быть несколько строк)
SELECT * FROM users WHERE age > 90;
DELETE FROM users WHERE age > 90;
-- DELETE 1   <- остался только "Временный"
SELECT * FROM users ORDER BY id;

-- 3) Удалить то, чего нет - тоже не ошибка
DELETE FROM users WHERE id = 9999;
-- DELETE 0

-- 4) Номер (id) удалённой строки не используется повторно
INSERT INTO users (name, email, age) VALUES ('Новый', 'new@mail.tj', 22);
SELECT id, name FROM users ORDER BY id;
--    В id появилась "дырка": номера удалённых записей не возвращаются.


-- =====================================================================
-- Мягкое удаление - строка остаётся, но помечена как "неактивная"
-- (в реальных проектах удаляют так очень часто)
-- =====================================================================
UPDATE users SET is_active = false WHERE name = 'Новый';
SELECT * FROM users WHERE is_active = true;     -- обычный список: "Нового" не видно
SELECT * FROM users WHERE is_active = false;    -- а в базе он по-прежнему есть
--    Вернуть человека - тоже обычный UPDATE:
UPDATE users SET is_active = true WHERE name = 'Новый';


-- =====================================================================
-- ОПАСНО: забытый WHERE
-- Эти две команды ЗАКОММЕНТИРОВАНЫ нарочно: если выполнить файл целиком,
-- они бы изменили/удалили ВСЮ таблицу. На уроке копируем по одной,
-- проговаривая: "WHERE забыли - значит, затронуты ВСЕ строки".
-- После показа возвращаем базу: reset.sql
-- =====================================================================
-- UPDATE users SET age = 0;
-- SELECT * FROM users;        -- у всех возраст 0
--
-- DELETE FROM users;
-- SELECT * FROM users;        -- таблица пуста
--
-- Вернуть всё обратно (в другом терминале):
--   docker compose exec -T postgres psql -U postgres -d coursedb < reset.sql


-- =====================================================================
-- Дайте студентам самим написать запросы (после reset.sql):
--   1) Поменять возраст Нигины на 28.
--   2) Сделать всех, кто моложе 26, неактивными.
--   3) Добавить пользователя "Диана", 24 года, и сразу удалить её.
--   4) Удалить всех неактивных. Сначала - SELECT, который покажет, кого удалим.
-- =====================================================================
