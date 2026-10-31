-- id сессии — случайный UUID, его генерирует приложение. Выход из сессии — удаление строки.
CREATE TABLE sessions (
    id          UUID        PRIMARY KEY,
    user_id     BIGINT      NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    device_type TEXT        NOT NULL DEFAULT 'web' CHECK (device_type IN ('web', 'mobile')),
    ip          INET,
    user_agent  TEXT        NOT NULL DEFAULT '',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    expires_at  TIMESTAMPTZ NOT NULL
);

CREATE INDEX idx_sessions_user_id ON sessions (user_id, created_at DESC);
CREATE INDEX idx_sessions_expires_at ON sessions (expires_at);
