-- Шаг 1. Создаём таблицу и наполняем её (выполняем в psql, команду за командой)
--
-- Подключение:
--   psql -h localhost -U postgres -d coursedb
-- Полезно: \dt - список таблиц, \d users - структура таблицы, \q - выход

CREATE TABLE users (
    id SERIAL PRIMARY KEY,                             -- уникальный номер, растёт сам
    name TEXT NOT NULL,                                -- обязательное поле
    email TEXT UNIQUE NOT NULL,                        -- не повторяется
    age INT CHECK (age >= 0 AND age <= 150),           -- проверка значения
    is_active BOOLEAN DEFAULT true,                    -- значение по умолчанию
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()  -- заполняет сама база, вместе с часовым поясом
);

INSERT INTO users (name, email, age) VALUES
    ('Али', 'ali@mail.tj', 25),
    ('Вера', 'vera@mail.tj', 30),
    ('Тимур', 'timur@mail.tj', 19),
    ('Нигина', 'nigina@mail.tj', 27);

SELECT * FROM users;


-- 1) email должен быть уникальным
INSERT INTO users (name, email, age) VALUES ('Дубль', 'ali@mail.tj', 20);
-- ERROR: duplicate key value violates unique constraint "users_email_key"

-- 2) возраст не может быть отрицательным
INSERT INTO users (name, email, age) VALUES ('Плохой', 'bad@mail.tj', -5);
-- ERROR: new row for relation "users" violates check constraint "users_age_check"
