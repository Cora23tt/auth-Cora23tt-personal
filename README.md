
# 🛡 Auth Service — практика

Простой микросервис авторизации с JWT и refresh токенами на Go с использованием PostgreSQL.

---

## 🚀 Описание

Реализовано:

- ✅ Регистрация по номеру телефона (`/register`)
- ✅ Логин с генерацией access и refresh токенов (`/login`)
- ✅ Обновление access токена по refresh (`/refresh`)
- ✅ Защищённый маршрут для получения информации (`/me`)
- ✅ PostgreSQL + pgxpool

---

## 🌐 Домен для проверки

> 🔗 После деплоя на Railway укажите здесь свой домен:

```
https://<your-subdomain>.up.railway.app
```

---

## ⚙️ Переменные окружения (.env)

```env
JWT_SECRET=your_jwt_secret
DATABASE_URL=postgres://user:password@host:5432/dbname
```

---

## ▶️ Запуск локально

```bash
go run main.go
```

---

## 🧪 Примеры curl-запросов

### Регистрация

```bash
curl -X POST https://<your-domain>/register \
  -H "Content-Type: application/json" \
  -d '{"phone":"+998901234567","password":"secret"}'
```

### Логин

```bash
curl -X POST https://<your-domain>/login \
  -H "Content-Type: application/json" \
  -d '{"phone":"+998901234567","password":"secret"}'
```

### Получить профиль

```bash
curl -X GET https://<your-domain>/me \
  -H "Authorization: Bearer <access_token>"
```

---

## 🗃 Используемые технологии

- Go + Gin
- PostgreSQL + pgxpool
- JWT: `github.com/golang-jwt/jwt/v5`
- UUID: `github.com/google/uuid`
- Пароли: `golang.org/x/crypto/bcrypt`
- Railway для деплоя

---

## 📦 Dockerfile (если нужен)

```Dockerfile
FROM golang:1.22-alpine
WORKDIR /app
COPY . .
RUN go build -o main .
CMD ["./main"]
```

---
