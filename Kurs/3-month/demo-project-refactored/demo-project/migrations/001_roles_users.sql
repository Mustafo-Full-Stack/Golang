CREATE TABLE roles (
    id          BIGINT      GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name        TEXT        NOT NULL,
    code        TEXT        NOT NULL UNIQUE,
    description TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- id выдаются по порядку вставки: 1, 2, 3...
-- Порядок строк должен совпадать с константами в internal/models/role.go.
-- Самостоятельно зарегистрироваться можно только как user, остальные роли назначает администратор.
INSERT INTO roles (name, code, description) VALUES
    ('Пользователь',  'user',      'Создаёт, редактирует и удаляет свои посты'),                       -- 1
    ('Администратор', 'admin',     'Полный доступ: пользователи, роли, настройки'),                    -- 2
    ('Аудитор',       'auditor',   'Только чтение: просмотр данных и истории действий для проверок'),  -- 3
    ('Модератор',     'moderator', 'Проверяет и удаляет посты, нарушающие правила'),                   -- 4
    ('Поддержка',     'support',   'Отвечает на обращения пользователей, видит их профили'),           -- 5
    ('Работодатель',  'employer',  'Публикует вакансии и рассматривает отклики'),                      -- 6
    ('Соискатель',    'applicant', 'Ведёт резюме и откликается на вакансии');                          -- 7

CREATE TABLE users (
    id            BIGINT      GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    email         TEXT        NOT NULL UNIQUE,
    password_hash TEXT        NOT NULL,
    name          TEXT        NOT NULL,
    role_id       BIGINT      NOT NULL REFERENCES roles(id),
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);
