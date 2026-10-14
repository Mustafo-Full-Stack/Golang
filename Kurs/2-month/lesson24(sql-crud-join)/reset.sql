-- Возвращает базу в состояние "начало урока": таблица users и 4 пользователя.
-- Остальные таблицы урока (tags, orders) мы создаём по ходу урока.
-- Запускать перед уроком и после любых экспериментов.
--
-- Docker Compose (из корня репозитория, путь к файлу - от корня):
--   docker compose exec -T postgres psql -U postgres -d coursedb < "lesson24(sql-crud-join)/reset.sql"
-- Локальный PostgreSQL:
--   psql -h localhost -U postgres -d coursedb -f reset.sql

DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS demo_prices;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    age INT CHECK (age >= 0 AND age <= 150),
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

INSERT INTO users (name, email, age) VALUES
    ('Али', 'ali@mail.tj', 25),
    ('Вера', 'vera@mail.tj', 30),
    ('Тимур', 'timur@mail.tj', 19),
    ('Нигина', 'nigina@mail.tj', 27);

SELECT * FROM users;
