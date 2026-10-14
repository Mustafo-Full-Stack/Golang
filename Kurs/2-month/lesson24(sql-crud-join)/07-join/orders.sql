-- Шаг 7. Вторая таблица, связанная с users (выполняем в psql)

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,   -- внешний ключ
    product TEXT NOT NULL,
    amount NUMERIC(10,2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

INSERT INTO orders (user_id, product, amount) VALUES
    (1, 'Ноутбук', 1200.00),
    (1, 'Мышь', 25.50),
    (2, 'Монитор', 300.00),
    (3, 'Клавиатура', 45.00);

SELECT * FROM orders;

-- База защищает целостность: пользователя 999 не существует
INSERT INTO orders (user_id, product, amount) VALUES (999, 'Товар', 10);
-- ERROR: insert or update on table "orders" violates foreign key constraint "orders_user_id_fkey"
