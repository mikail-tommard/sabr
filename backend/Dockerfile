FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

FROM alpine:3.20

WORKDIR /app

RUN apk add --no-cache ca-certificates bash make

COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /app/api ./api
COPY . .

EXPOSE 8080

CMD ["./api"]
