package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kenny-mwendwa/go-restapi-crud/config"
)

func ConnectDB() (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DBConfig.Host, config.DBConfig.Port, config.DBConfig.User, config.DBConfig.Password, config.DBConfig.Name)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Println("Failed to ping database:", err)
		return nil, err
	}

	return db, nil
}
