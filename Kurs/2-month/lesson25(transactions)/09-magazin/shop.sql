-- Шаг 9. Таблицы магазина (выполняем в psql перед запуском main.go).
-- Файл можно запускать повторно: он сам пересоздаёт таблицы.

DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS products;

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL,
    quantity INT NOT NULL CHECK (quantity >= 0)     -- остаток на складе
);

-- Покупатель - это владелец счёта из таблицы accounts
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    account_id INT NOT NULL REFERENCES accounts(id),
    status TEXT NOT NULL,
    total NUMERIC(12,2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL REFERENCES orders(id),
    product_id INT NOT NULL REFERENCES products(id),
    quantity INT NOT NULL,
    price NUMERIC(10,2) NOT NULL                    -- цена на момент покупки
);

INSERT INTO products (name, price, quantity) VALUES
    ('Ноутбук', 800.00, 5),
    ('Мышь', 25.50, 20),
    ('Клавиатура', 45.00, 3),
    ('Монитор', 300.00, 2);

-- Пополним счёт Али, чтобы хватило на покупки
UPDATE accounts SET balance = 2000.00 WHERE id = 1;

SELECT * FROM products;
