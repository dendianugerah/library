# Library Management API

A REST API for managing library books, borrowers, and loans.

## Getting Started

### Prerequisites
- Go 1.16 or higher
- Git
- PostgreSQL 14 or higher

### Database Setup

1. Create a PostgreSQL database
```sql
CREATE DATABASE library_db;
```

2. Update database connection in `db/postgres.go`:
```go
// Change this line with your database credentials
dsn := "postgresql://username:password@localhost:5432/library_db?sslmode=disable"
```

### Quick Start

1. Clone and install dependencies
```bash
git clone https://github.com/yourusername/library_api.git
cd library_api
go mod download
```

2. Run the application
```bash
go run main.go
```
The server will start at `http://localhost:8080` and will automatically:
- Connect to the database
- Create required tables
- Set up the API endpoints

## API Documentation

### View API Docs
Access the Swagger UI at: `http://localhost:8080/swagger/index.html`

### Core Endpoints

#### Books
- `POST /books` - Create a new book
- `GET /books` - List all books

#### Borrowers
- `POST /borrowers` - Register new borrower
- `GET /borrowers` - List all borrowers

#### Loans
- `POST /loans` - Create a new loan
- `GET /loans` - List all loans (can filter by status)
- `PUT /loans/{id}/return` - Return a book

Example: Get late returns
```bash
curl http://localhost:8080/loans?status=late
```

Example: Get on-time returns
```bash
curl http://localhost:8080/loans?status=ontime
```

### Business Rules

- One book per loan
- Maximum loan duration: 30 days
- No new loans while having active loans
- Must have available stock to borrow
- Books must have unique ISBN
