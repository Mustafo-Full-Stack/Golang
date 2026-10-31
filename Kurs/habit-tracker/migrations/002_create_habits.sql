CREATE TABLE IF NOT EXISTS habits (
    id         BIGSERIAL PRIMARY KEY,
    user_id    BIGINT      NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    name       TEXT        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS habit_checks (
    id       BIGSERIAL PRIMARY KEY,
    habit_id BIGINT NOT NULL REFERENCES habits (id) ON DELETE CASCADE,
    date     DATE   NOT NULL,
    UNIQUE (habit_id, date)
);
