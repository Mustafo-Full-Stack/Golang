COMPOSE := docker compose -f docker-compose.infra.yaml

.PHONY: run test infra-up infra-down infra-reset

# Запустить приложение (сначала поднимите инфраструктуру: make infra-up)
run:
	go run ./cmd/api

# Проверить код перед коммитом
test:
	gofmt -l internal pkg cmd
	go vet ./...
	go test ./...

# Поднять Postgres, Redis и MinIO и дождаться, пока они будут готовы
infra-up:
	$(COMPOSE) up -d --wait

# Остановить. Данные базы сохраняются
infra-down:
	$(COMPOSE) down

# Удалить всё вместе с данными и поднять заново: миграции применятся с нуля
infra-reset:
	$(COMPOSE) down -v
	$(COMPOSE) up -d --wait
