package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kenny-mwendwa/go-restapi-crud/config"
)

func ConnectDB() (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return db, nil
}
