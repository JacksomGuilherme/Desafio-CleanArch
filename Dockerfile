FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ordersystem ./cmd/ordersystem

FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/ordersystem .

COPY --from=builder /app/.env* .

COPY --from=builder /app/internal/infra/database/migrations ./internal/infra/database/migrations

EXPOSE 8000
EXPOSE 8080
EXPOSE 50051

CMD ["./ordersystem"]
