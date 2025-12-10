# Chirpy — A Minimal Twitter-like Microservice in Go

Chirpy is a lightweight microservice written in Go that provides:

- User registration & login with **JWT authentication**
- Posting and reading short messages (“chirps”)
- Persistent storage using **PostgreSQL**
- **Goose** for versioned database migrations
- **sqlc** for type-safe database access
- A minimal built-in web UI
- Health/readiness endpoints

---

## 🚀 Features

### 🔐 Authentication
- Secure password hashing
- JWT-based access tokens
- Auth middleware
- Token expiry & verification

### 🐦 Chirps
- Create chirps
- List all chirps
- List chirps by user

### 🗄️ Database Layer
- PostgreSQL
- Goose migrations
- SQLC generated code

### 🩺 Infra
- `/api/healthz` readiness
- `/api/reset` dev-only reset route

---

## 📁 Project Structure

```
.
├── assets/              # Static frontend files
├── sql/
│   ├── migrations/      # Goose migrations (0001_init.sql, etc.)
│   └── queries/         # SQLC query definitions
├── chirps.go            # Chirp handlers
├── users.go             # User registration logic
├── login.go             # Login & JWT generation
├── json.go              # JSON helper functions
├── readiness.go         # /api/healthz endpoint
├── reset.go             # Dev reset endpoint
├── index.html           # Simple UI page
├── sqlc.yaml            # SQLC config
├── go.mod
└── go.sum
```

---

## 🛠️ Requirements

- Go 1.20+
- PostgreSQL 14+
- Goose (`go install github.com/pressly/goose/v3/cmd/goose@latest`)
- sqlc (`go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`)

---

## 🔧 Environment Variables

```bash
export DATABASE_URL="postgres://user:password@localhost:5432/chirpy?sslmode=disable"
export JWT_SECRET="your_super_secret_key"
export PORT=8080
```

---

## 🗄️ Database Setup

### Install Goose

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### Create Database

```bash
createdb chirpy
```

### Run Migrations

```bash
goose -dir sql/migrations postgres "$DATABASE_URL" up
```

### (Optional) Generate SQLC Code

```bash
sqlc generate
```

---

## ▶️ Running the Server

### Development

```bash
go run .
```

### Production

```bash
go build -o chirpy .
./chirpy
```

---

## 🧪 Example API Requests

### Signup

```bash
curl -X POST http://localhost:8080/api/users   -H "Content-Type: application/json"   -d '{"email":"test@example.com","password":"secret"}'
```

### Login

```bash
curl -X POST http://localhost:8080/api/login   -H "Content-Type: application/json"   -d '{"email":"test@example.com","password":"secret"}'
```

### Create Chirp

```bash
curl -X POST http://localhost:8080/api/chirps   -H "Authorization: Bearer <TOKEN>"   -H "Content-Type: application/json"   -d '{"body":"hello world"}'
```

### List Chirps

```bash
curl http://localhost:8080/api/chirps
```

---

## 📜 License
MIT
