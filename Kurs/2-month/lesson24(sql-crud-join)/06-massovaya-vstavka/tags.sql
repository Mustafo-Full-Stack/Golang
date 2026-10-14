-- Таблица для опыта с массовой вставкой (выполняем в psql перед запуском main.go)
CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);
