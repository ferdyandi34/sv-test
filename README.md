# Article Service + Frontend

This project consists of:

- A Go backend (REST API with GORM + MySQL)
- A Vite + React frontend (styled with Chakra UI)

---

## ⚙️ Setup

### 1. Environment Variables

Create a `.env` file in the backend root:

```env
DB_USER=root
DB_PASS=secret
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=mydb

```

### 2. Run MySQL Service

Download and install MySQL Workbench

### 3. Run Backend with Air

install air for live reload:
go install github.com/cosmtrek/air@latest

run command

```air

```

### 4. Run Frontend

Navigate into frontend client folder

```cd client

```

Install dependencies

```npm install

```

Build the project

```npm run build

```

Start dev server

```npm run dev

```
