-- Шаг 1. Готовим таблицы счетов (выполняем в psql)
--   psql -h localhost -U postgres -d coursedb

CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    owner TEXT NOT NULL,
    balance NUMERIC(12,2) NOT NULL CHECK (balance >= 0)   -- база сама не даст уйти в минус
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

-- Проверим CHECK: попробуем уйти в минус
UPDATE accounts SET balance = -1 WHERE id = 1;
-- ERROR: new row for relation "accounts" violates check constraint "accounts_balance_check"
