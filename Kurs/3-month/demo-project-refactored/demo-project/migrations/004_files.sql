CREATE TABLE files (
    id            BIGINT      GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id       BIGINT      NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    -- NULL — файл загружен, но ещё не прикреплён к посту.
    post_id       BIGINT      REFERENCES posts(id) ON DELETE CASCADE,
    object_key    TEXT        NOT NULL UNIQUE,
    original_name TEXT        NOT NULL,
    content_type  TEXT        NOT NULL,
    size_bytes    BIGINT      NOT NULL CHECK (size_bytes > 0),
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_files_post_id ON files (post_id);
CREATE INDEX idx_files_unattached ON files (created_at) WHERE post_id IS NULL;
