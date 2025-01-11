# GO Frameworks RESTful CRUD Application

This is a simple CRUD (Create, Read, Update, Delete) application written in Go, showcasing different routers and frameworks. The application allows you to perform operation on user data using various web frameworks and routers, including Chi, Echo, Gin, HttpRouter, Mux, and the standard net/http package.

## Project Overview

The project follows the standard Go project layout and naming conventions, providing a clean and modular structure:

- `cmd/`: Contains the main application entry point (`main.go`). This is where the main application logic resides.
- `internals/`: Contains internal packages and modules that are specific to this application. These packages are not intended to be imported by external packages.
  - `config/`: Handles application configuration, such as database configuration for the apis.
  - `database/`: Contains database related packages.
    - `connection.go`: Creates a pooled database connection to a postgres database using [pgx](https://github.com/jackc/pgx).
    - `db.go`, `models.go` & `users.sql.go`: Contains [sqlc](https://docs.sqlc.dev/en/latest/index.html) generated type safe GO code generated from `schema.sql` and `query.sql` files.
  - `migrate/`: Contains script `main.go` for database migration using golang-migrate.
  - `routers/`: Contains router implementations (`chi_router.go`, `echo_router.go`, etc.).
  - `handlers/`: Contains handler functions for each CRUD operation (`chi_handler.go`, `echo_handler.go`, etc.).
  - `sql/`: Contains `schema` and `queries` folders with sql files for generating type safe GO code from the compiled sql using sqlc.
- `sqlc.yaml`: This is the configuration file used for working with [sqlc](https://docs.sqlc.dev/en/latest/index.html).

## Getting Started

### Prerequisites

Ensure you have the following installed on your machine:

- [Go](https://golang.org/dl/): The application is developed in Go. Install the latest version from the official website.
- PostgreSQL Database: The application uses PostgreSQL as the database. Ensure it is installed and running.
- [sqlc] (https://docs.sqlc.dev/en/latest/): sqlc generates fully type-safe idiomatic Go code from SQL.
- [goose] (https://pressly.github.io/goose/): is a database migration tool. Manage your database schema by creating incremental SQL changes and/or Go functions.

### Environment Configuration

Create a `.env` file in the root of the project with the following content:

```env
DATABASE_URL="postgresql://DB_USER:DB_PASSWORD@DB_HOST:DB_PORT/DB_NAME"

```

### Running Migrations

To initialize the database schema, navigate to the schema directory using:

```bash
cd internal/sql/schema
```

and run the following command:

```bash
goose postgres postgresql://username:password@localhost:5432/db_name up
```

### Running the application

To build and run the application use the following command

```bash
make run
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
