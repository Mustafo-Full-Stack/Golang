# crud-api

## Запуск in-memory

```
go run ./cmd/api
```

По умолчанию `IS_REPO_IN_MEMORY=true` и `IS_TOKEN_IN_MEMORY=true`: пользователи и токены хранятся в памяти.

## Запуск с Postgres

1. Скопировать `.env.example` в `.env`, выставить `IS_REPO_IN_MEMORY=false` и заполнить `POSTGRES_*`.
2. Применить миграцию:

```
psql "$DSN" -f migrations/001_create_users.sql
```

3. `go run ./cmd/api`

## Токены в Redis

В `.env` выставить `IS_TOKEN_IN_MEMORY=false` и заполнить `REDIS_*`. Токены хранятся под ключами `token:<uuid>`.

Время жизни токена задаётся `TOKEN_TTL` в формате Go duration (`30m`, `24h`), по умолчанию `24h`. Работает и для in-memory, и для Redis.

## Логирование

По умолчанию `IS_SLOG_ENABLED=false`, используется обычный `log`. При `IS_SLOG_ENABLED=true` логи идут через `slog` с уровнем `LOG_LEVEL`:
- `ENVIRONMENT=local`: текстом в файл `LOG_FILE`;
- иначе: JSON в stdout.

## Пример

```
go run ./cmd/example
```
