# 💸 Expense Tracker API

A RESTful API built with **Go** and **PostgreSQL** for tracking personal expenses. Features JWT authentication, bcrypt password hashing, and a clean, modular project structure.

---

## 🗂️ Project Structure

```
ExpenseTracker/
├── config/
│   └── db.go           # PostgreSQL connection setup
├── handlers/
│   ├── auth.go         # Register & Login handlers
│   └── expense.go      # CRUD handlers for expenses
├── middleware/
│   └── auth.go         # JWT authentication middleware
├── models/
│   ├── user.go         # User struct & DB queries
│   └── expense.go      # Expense struct & DB queries
├── router/
│   └── router.go       # Route definitions
├── .env                # Environment variables (not committed)
├── .gitignore
├── go.mod
├── go.sum
└── main.go             # Entry point
```

---

## ⚙️ Tech Stack

| Technology | Purpose |
|---|---|
| Go (net/http) | REST API server |
| PostgreSQL | Database |
| JWT (golang-jwt/jwt) | Authentication |
| bcrypt | Password hashing |
| godotenv | Environment variable loading |

---

## 🚀 Getting Started

### Prerequisites
- Go 1.21+
- PostgreSQL

### 1. Clone the repository
```bash
git clone https://github.com/Saugat-Tamang17/ExpenseTracker.git
cd ExpenseTracker
```

### 2. Install dependencies
```bash
go get github.com/lib/pq
go get github.com/golang-jwt/jwt/v5
go get github.com/joho/godotenv
go get golang.org/x/crypto/bcrypt
```

### 3. Setup PostgreSQL
```sql
CREATE DATABASE expensetracker;

\c expensetracker

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE expenses (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(200) NOT NULL,
    amount NUMERIC(10, 2) NOT NULL,
    category VARCHAR(100),
    note TEXT,
    date DATE DEFAULT CURRENT_DATE,
    created_at TIMESTAMP DEFAULT NOW()
);
```

### 4. Create your `.env` file
```env
DB_URL=postgres://postgres:yourpassword@localhost:5432/expensetracker?sslmode=disable
JWT_SECRET=your_super_secret_key
PORT=9090
```

### 5. Run the server
```bash
go run main.go
```

Server will start at `http://localhost:9090` ✅

---

## 📡 API Endpoints

### Auth (Public)

| Method | Endpoint | Description |
|---|---|---|
| POST | `/register` | Register a new user |
| POST | `/login` | Login and receive JWT token |

#### Register
```json
POST /register
{
  "username": "saugat",
  "email": "saugat@example.com",
  "password": "yourpassword"
}
```

#### Login
```json
POST /login
{
  "email": "saugat@example.com",
  "password": "yourpassword"
}
```
Returns:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIs..."
}
```

---

### Expenses (Protected)
> All expense routes require the JWT token in the Authorization header:
> `Authorization: Bearer <your_token>`

| Method | Endpoint | Description |
|---|---|---|
| GET | `/expenses` | Get all expenses for logged in user |
| POST | `/expenses` | Create a new expense |
| PUT | `/expenses/:id` | Update an expense |
| DELETE | `/expenses/:id` | Delete an expense |

#### Create Expense
```json
POST /expenses
{
  "title": "Grocery",
  "amount": 45.50,
  "category": "Food",
  "note": "Weekly groceries",
  "date": "2026-03-12T00:00:00Z"
}
```

---

## 🔐 Authentication Flow

1. Register or Login to receive a JWT token
2. Include the token in all protected requests as:
   ```
   Authorization: Bearer <token>
   ```
3. Token expires after **24 hours**

---

## 🧱 Architecture

The project follows a clean separation of concerns:

- **`config/`** — Initializes and exposes the DB connection
- **`models/`** — All SQL queries and data structs live here, handlers never write raw SQL
- **`handlers/`** — Reads HTTP requests, calls models, returns JSON responses
- **`middleware/`** — Intercepts protected routes and verifies JWT before the handler runs
- **`router/`** — Single source of truth for all routes and which ones are protected

---

## 👤 Author

**Saugat Tamang**
- GitHub: [@Saugat-Tamang17](https://github.com/Saugat-Tamang17)
- Student ID: 024BSCIT036
