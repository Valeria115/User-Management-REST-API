# User Management REST API

Простой REST API на Go для управления пользователями с PostgreSQL.

## Возможности

- **GET /users** — получить всех пользователей
- **GET /users/{id}** — получить пользователя по ID
- **POST /users** — создать нового пользователя
- **PUT /users/{id}** — обновить пользователя
- **DELETE /users/{id}** — удалить пользователя

## Требования

- Go 1.27+
- PostgreSQL 15+
- Docker & Docker Compose (опционально)

## Быстрый старт с Docker

```bash
docker-compose up
```

API будет доступен на `http://localhost:8080`

## Локальный запуск

1. **Подключитесь к PostgreSQL:**

```bash
psql -U postgres
```

2. **Создайте БД:**

```sql
CREATE DATABASE go_practice;
```

3. **Создайте таблицу:**

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE
);
```

4. **Установите зависимости:**

```bash
go mod download
```

5. **Запустите приложение:**

```bash
go run main.go
```

## Примеры запросов

```bash
# Создать пользователя
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Alice", "email": "alice@example.com"}'

# Получить всех пользователей
curl http://localhost:8080/users

# Получить пользователя по ID
curl http://localhost:8080/users/1

# Обновить пользователя
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Alice Updated", "email": "alice.updated@example.com"}'

# Удалить пользователя
curl -X DELETE http://localhost:8080/users/1
```

## Тестирование

```bash
go test -v
```

## Структура проекта

```
.
├── main.go              # API endpoints
├── main_test.go         # Setup для тестов
├── *_test.go            # Тесты для каждого handler'а
├── go.mod               # Go модули
├── Dockerfile           # Конфигурация Docker
├── docker-compose.yml   # Docker Compose конфигурация
├── init.sql            # SQL для инициализации БД
└── README.md           # Этот файл
```

## Переменные окружения

При локальном запуске можно переопределить:

```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=go_practice
```

## Примечания

- Все email'ы должны быть уникальными
- ID генерируются автоматически
- Все ошибки возвращаются в JSON формате
