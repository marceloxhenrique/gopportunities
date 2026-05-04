# Go Opportunities API

## A REST API built with **Go** and **Gin** for managing job opportunities. This project focuses on building and testing CRUD operations using GORM with a SQL database, along with an in-memory repository for testing, and clear separation between handlers and data access logic.

## Features

- Create a job opening
- List all openings
- Get opening by ID
- Update an opening
- Delete an opening
- In-memory repository (no database required)
- Unit tests for handlers

---

## Project Structure

```
.
├── config/         # App configuration
├── docs/           # Documentation
├── handler/        # HTTP handlers
├── repository/     # Data access layer
├── router/         # Route definitions
├── schemas/        # Data models
└── main.go         # Entry point
```

---

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/marceloxhenrique/gopportunities.git
cd gopportunities
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Run the application

```bash
go run main.go
```

Server will start on:

```
http://localhost:8080
```

## Running Tests

```bash
go test -v ./...
```

---

## Concepts Demonstrated

- REST API design
- Separation of concerns (handler vs repository)
- Dependency injection
- JSON serialization/deserialization
- Unit testing with `httptest`

---

## Future Improvements

- Add validation
- Add authentication
- Pagination for listings

---
