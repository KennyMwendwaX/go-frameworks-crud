package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading DATABASE_URL from .env file")
	}

	connStr := os.Getenv("DATABASE_URL")

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	defer db.Close()

	// Create a new migrate instance
	dbInstance, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to create migrate driver: %v", err)
	}

	// Run migrations
	m, err := migrate.NewWithDatabaseInstance("file://internals/sql", "postgres", dbInstance)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations applied successfully")
}
