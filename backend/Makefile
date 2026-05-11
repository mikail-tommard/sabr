APP=./cmd/api
COMPOSE=docker compose -f ../docker-compose.yml
DB_DSN ?= postgres://sabr:sabr@postgres:5432/sabr?sslmode=disable

.PHONY: run build tidy fmt migrate-up migrate-down migrate-status

run:
	@go run $(APP)

build:
	@go build $(APP)

tidy:
	@go mod tidy

fmt:
	@go fmt ./...

migrate-up:
	@$(COMPOSE) run --rm -e DB_DSN='$(DB_DSN)' backend sh -lc 'goose -dir ./migrations postgres "$$DB_DSN" up'

migrate-down:
	@$(COMPOSE) run --rm -e DB_DSN='$(DB_DSN)' backend sh -lc 'goose -dir ./migrations postgres "$$DB_DSN" down'

migrate-status:
	@$(COMPOSE) run --rm -e DB_DSN='$(DB_DSN)' backend sh -lc 'goose -dir ./migrations postgres "$$DB_DSN" status'
