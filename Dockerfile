# --- Этап 1: сборка бинарного файла ---
FROM golang:1.26-alpine AS builder
LABEL authors="Dreamod"

WORKDIR /build

# ставит зависимости
COPY go.mod go.sum ./
RUN go mod download

# копирует файлы проекта
COPY . .

# сборка бинарника приложения
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o app ./cmd/main.go

# --- Этап 2: финальный образ (минимальный) ---
FROM alpine:3.19

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /build/app .

EXPOSE 8081

CMD ["./app"]
