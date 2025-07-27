# Library REST API (Go + PostgreSQL)

A minimal **REST API** for managing books in a public library, built with **Go**, **PostgreSQL**, and **Docker**.  
It supports **CRUD operations** on books, has **token-based authentication**, **structured logging**, **unit tests**, and **database migrations** using `golang-migrate`.

---

## **Features**
- RESTful API with [Gin](https://github.com/gin-gonic/gin).
- PostgreSQL with schema migrations using [golang-migrate](https://github.com/golang-migrate/migrate).
- Token-based authentication middleware.
- Structured request logging.
- Health check endpoint (`/health`).
- Unit tests with `testify`.
- Use Dockerfile and Docker Compose for easy setup.
- 

---

## **API Endpoints**

### **Public Endpoint**
| Method | Endpoint     | Description        |
|--------|--------------|--------------------|
| GET    | `/health`    | Health check (no token required) |

### **Protected Endpoints** (require token in `Authorization` header)
| Method | Endpoint       | Description           |
|--------|----------------|-----------------------|
| GET    | `/books`       | List all books        |
| GET    | `/books/:id`   | Get book by ID        |
| POST   | `/books`       | Create a new book     |
| PUT    | `/books/:id`   | Update existing book  |
| DELETE | `/books/:id`   | Delete book by ID     |

---

## **cURL Examples**

Assume `API_TOKEN=my-secret-token` (defined in `docker-compose.yml`).

### **Health Check**
```bash
curl http://localhost:8080/health
```

### **List All Books**
```bash
curl http://localhost:8080/books \
  -H "Authorization: my-secret-token"
```

### **Create a Book**
```bash
curl -X POST http://localhost:8080/books \
  -H "Authorization: my-secret-token" \
  -H "Content-Type: application/json" \
  -d '{"title":"The Hobbit","author":"J.R.R. Tolkien","published_at":"1937"}'
```

### **Get a Book by ID**
```bash
curl http://localhost:8080/books/1 \
  -H "Authorization: my-secret-token"
```

### **Update a Book**
```bash
curl -X PUT http://localhost:8080/books/1 \
  -H "Authorization: my-secret-token" \
  -H "Content-Type: application/json" \
  -d '{"title":"The Hobbit - Revised","author":"J.R.R. Tolkien","published_at":"1937"}'
```

### **Delete a Book**
```bash
curl -X DELETE http://localhost:8080/books/1 \
  -H "Authorization: my-secret-token"
```

---

## **Setup**

### **1. Prerequisites**
- [Go 1.22+](https://go.dev/dl/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [golang-migrate](https://github.com/golang-migrate/migrate) (for manual migrations).

---

### **2. Clone & Build**
```bash
git clone https://github.com/your-username/library-api.git
cd library-api
```

---

### **3. Start the API and Database**
```bash
make run
```
This will:
- Start a **PostgreSQL** container.
- Build and run the **Go REST API**.
- API available at: `http://localhost:8080`.

---

## **Database Migrations**

### **Create a Migration**
```bash
migrate create -ext sql -dir db/migrations -seq add_books_table
```

### **Apply Migrations**
```bash
make migrate-up
```

### **Rollback**
```bash
make migrate-down
```

---

## **Testing**

### **Run Unit Tests**
```bash
make test
```
- Tests cover all `books` handler endpoints and middleware.
- Uses a mock repository implementing `BookRepository` for clean, isolated tests.

---

## **Authentication**

All routes (except `/health`) require a token:
```bash
-H "Authorization: my-secret-token"
```

**Set token in `.yml`**:
```yaml
  environment:
    API_TOKEN: my-secret-token
```

---

## **Logging**

- Every request is logged (method, path, IP, response status, and latency).
- CRUD operations log actions and errors.

---

## **Directory Structure**
```
library-api/
├── cmd/server/main.go          # Entry point
├── config/config.go            # DB and environment config
├── db/migrations/              # SQL migration files
├── internal/books/             # Book model, handler, repository
├── internal/database/db.go     # DB initialization
├── internal/router/router.go   # Routes and middleware
├── internal/middleware/        # Request logging & token auth
├── Dockerfile
├── docker-compose.yml
├── Makefile
├── go.mod / go.sum
└── README.md
```

---

## **Next Steps**
- Add GitHub Actions for automated testing.
- Extend the API with pagination, search, and sorting.
- Implement JWT-based authentication (optional).