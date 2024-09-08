FROM golang:1.21-alpine

# Ставим рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum файлы
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем весь остальной код в контейнер
COPY . .

# Копируем файл serviceAccountKey.json в контейнер
COPY serviceAccountKey.json /app/serviceAccountKey.json

# Проверка сборки
RUN go build -o main cmd/main.go

# Запускаем приложение
CMD ["./main"]