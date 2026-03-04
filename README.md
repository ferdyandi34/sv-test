# Article Service + Frontend

This project consists of:

- A Go backend (REST API with GORM + MySQL) => Tested in postman
- A Vite + React frontend (styled with Chakra UI) => On debugging, feature coded: list and create

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

```
air

```

### 4. Run Frontend

Navigate into frontend client folder

```
cd client

```

Install dependencies

```
npm install

```

Build the project

```
npm run build

```

Start dev server

```
npm run dev

```

### CURL for Backend Testing

- Create Article

```
curl --location 'http://localhost:5000/articles/create' \
--header 'Content-Type: text/plain' \
--data '{
    "title": "The Fishes of the sea",
    "content": "tidebringer",
    "category": "novel",
    "status": "draft"
}'

```

- List Article

```
curl --location 'http://localhost:5000/articles?limit=10&offset=0'

```

- Get Article by Id

```
curl --location 'http://localhost:5000/articles/get?id=2'

```

- Update Article

```
curl --location --request PUT 'http://localhost:5000/articles/update?id=1' \
--header 'Content-Type: text/plain' \
--data '{
    "title": "The Ultimate Whale combat",
    "content": "ultimate whale",
    "category": "novel",
    "status": "publish"
}'

```

- Delete Article

```
curl --location --request DELETE 'http://localhost:5000/articles/delete?id=1'

```
