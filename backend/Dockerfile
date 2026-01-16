# Build stage
FROM golang:1.24.0-alpine AS builder

WORKDIR /app

# Копируем go mod файлы
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение из cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Копируем бинарник из builder
COPY --from=builder /app/main .

# Копируем конфиги если есть
COPY --from=builder /app/config ./config

# Создаем директорию для uploads
RUN mkdir -p /root/uploads

EXPOSE 8080

CMD ["./main"]