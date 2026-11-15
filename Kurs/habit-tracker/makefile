# Если файл `.env` существует — прочитать его (знак `-` спереди означает "не ругайся, если файла `.env` ещё нет").
-include .env
# Передать все прочитанные переменные из `.env` в запускаемые команды (`go run`, `docker compose`).
export
# Список команд-кнопок (чтобы `make` знал, что это названия действий, а не имена файлов).
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
# Посмотреть живые логи работающего контейнера приложения.
logs:
	docker compose logs -f app
# Автоматически красиво выровнять отступы во всех `.go` файлах проекта.
fmt:
	gofmt -w $$(find . -name '*.go' -not -path './.git/*')
# Проверить код анализатором Go на скрытые ошибки.
vet:
	go vet ./...
# Проверка перед коммитом: форматирование, vet и тесты.
test:
	test -z "$$(gofmt -l .)"
	go vet ./...
	go test ./...