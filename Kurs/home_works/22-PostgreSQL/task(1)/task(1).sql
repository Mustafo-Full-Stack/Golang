CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    price INT CHECK ( price >= 0 ),
    pages INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
)