-- Возвращает базу в состояние "после шага 1": счета Али, Веры и Тимура.
-- Запускать перед уроком, ПОСЛЕ шага 2 (там деньги "пропадают") и после любых экспериментов.
-- Таблицы магазина (products, orders, order_items) создаются на шаге 9 файлом shop.sql.
--
-- Docker Compose (из корня репозитория, путь к файлу - от корня):
--   docker compose exec -T postgres psql -U postgres -d coursedb < "lesson25(transactions)/reset.sql"
-- Локальный PostgreSQL:
--   psql -h localhost -U postgres -d coursedb -f reset.sql

DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS transfers;
DROP TABLE IF EXISTS accounts;

CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    owner TEXT NOT NULL,
    balance NUMERIC(12,2) NOT NULL CHECK (balance >= 0)
);

CREATE TABLE transfers (
    id SERIAL PRIMARY KEY,
    from_id INT REFERENCES accounts(id),
    to_id INT REFERENCES accounts(id),
    amount NUMERIC(12,2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

INSERT INTO accounts (owner, balance) VALUES
    ('Али', 1000.00),
    ('Вера', 500.00),
    ('Тимур', 250.00);

SELECT * FROM accounts;
