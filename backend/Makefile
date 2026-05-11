APP=./cmd/api

.PHONY: run build tidy

run:
	@go run $(APP)

build:
	@go build $(APP)

tidy:
	@go mod tidy

fmt:
	@go fmt ./...