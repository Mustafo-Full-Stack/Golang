-include .env
export

.PHONY: run db up down reset logs fmt vet test

# Запустить API локально (база должна работать: make db).
run:
	go run ./cmd/api

# Поднять только PostgreSQL и дождаться готовности.
db:
	docker compose up -d --wait db

# Поднять базу и приложение в Docker.
up:
	docker compose up --build -d

# Остановить контейнеры, данные базы сохраняются.
down:
	docker compose down

# Удалить контейнеры и данные базы.
reset:
	docker compose down -v

logs:
	docker compose logs -f app

fmt:
	gofmt -w $$(find . -name '*.go' -not -path './.git/*')

vet:
	go vet ./...

# Проверка перед коммитом: форматирование, vet и тесты.
test:
	test -z "$$(gofmt -l .)"
	go vet ./...
	go test ./...
