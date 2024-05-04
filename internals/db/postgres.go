package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func ConnectDB() (*pgx.Conn, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading DATABASE_URL from .env file")
	}

	connStr := os.Getenv("DATABASE_URL")

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	return conn, nil
}
