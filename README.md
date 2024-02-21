# Go RESTful CRUD Application

This is a simple CRUD (Create, Read, Update, Delete) application written in Go, showcasing different routers and frameworks. The application allows you to perform operation on user data using various web frameworks and routers, including Chi, Echo, Gin, HttpRouter, Mux, and the standard net/http package.

## Project Overview

The project follows the standard Go project layout and naming conventions, providing a clean and modular structure:

- `cmd/`: Contains the main application entry point (`main.go`). This is where the main application logic resides.
- `internal/`: Contains internal packages and modules that are specific to this application. These packages are not intended to be imported by external packages.
  - `config/`: Handles application configuration, such as database configuration (`db.go`).
  - `db/`: Handles database connectivity (`postgres.go`).
  - `models/`: Defines application data models (`user.go`).
  - `migrate/`: Contains scripts for database migrations (`main.go`).
  - `routers/`: Contains router implementations (`chi_router.go`, `echo_router.go`, etc.).
  - `handlers/`: Contains handler functions for each CRUD operation (`chi_handler.go`, `echo_handler.go`, etc.).

## Getting Started

### Prerequisites

Ensure you have the following installed on your machine:

- [Go](https://golang.org/dl/)
- PostgreSQL Database

### Requirements

- [Go](https://golang.org/dl/): The application is developed in Go. Install the latest version from the official website.

- PostgreSQL Database: The application uses PostgreSQL as the database. Ensure it is installed and running.

### Environment Configuration

Create a `.env` file in the root of the project with the following content:

```env
DB_HOST=your_database_host
DB_PORT=your_database_port
DB_USER=your_database_user
DB_PASSWORD=your_database_password
DB_NAME=your_database_name
```

### Running Migrations

To initialize the database schema, run the migration script:

```bash
go run internal/migrate/main.go
```

### Running the application

To initialize the database schema, run the migration script:

```bash
go run cmd/main.go
```

## Routers and Endpoints

### 1. Standard library: `net/http`

**Running at:** http://localhost:8000

- **Get All Users:**
  ```plaintext
  GET /users
  ```
- **Create User:**
  ```plaintext
  POST /users
  ```
- **Get User:**
  ```plaintext
  GET /users/:id
  ```
- **Update User:**
  ```plaintext
  PUT /users/:id
  ```
- **Delete User:**
  ```plaintext
  DELETE /users/:id
  ```

### 2. httprouter Router

**Running at:** http://localhost:8001

- **Get All Users:**
  ```plaintext
  GET /users
  ```
- **Create User:**
  ```plaintext
  POST /users
  ```
- **Get User:**
  ```plaintext
  GET /users/:id
  ```
- **Update User:**
  ```plaintext
  PUT /users/:id
  ```
- **Delete User:**
  ```plaintext
  DELETE /users/:id
  ```

### 3. Mux Router

**Running at:** http://localhost:8002

- **Get All Users:**
  ```plaintext
  GET /users
  ```
- **Create User:**
  ```plaintext
  POST /users
  ```
- **Get User:**
  ```plaintext
  GET /users/:id
  ```
- **Update User:**
  ```plaintext
  PUT /users/:id
  ```
- **Delete User:**
  ```plaintext
  DELETE /users/:id
  ```

### 4. Chi Router

**Running at:** http://localhost:8003

- **Get All Users:**
  ```plaintext
  GET /users
  ```
- **Create User:**
  ```plaintext
  POST /users
  ```
- **Get User:**
  ```plaintext
  GET /users/:id
  ```
- **Update User:**
  ```plaintext
  PUT /users/:id
  ```
- **Delete User:**
  ```plaintext
  DELETE /users/:id
  ```

### 5. Echo

**Running at:** http://localhost:8004

- **Get All Users:**
  ```plaintext
  GET /users
  ```
- **Create User:**
  ```plaintext
  POST /users
  ```
- **Get User:**
  ```plaintext
  GET /users/:id
  ```
- **Update User:**
  ```plaintext
  PUT /users/:id
  ```
- **Delete User:**
  ```plaintext
  DELETE /users/:id
  ```

### 6. Gin

**Running at:** http://localhost:8005

- **Get All Users:**
  ```plaintext
  GET /users
  ```
- **Create User:**
  ```plaintext
  POST /users
  ```
- **Get User:**
  ```plaintext
  GET /users/:id
  ```
- **Update User:**
  ```plaintext
  PUT /users/:id
  ```
- **Delete User:**
  ```plaintext
  DELETE /users/:id
  ```
