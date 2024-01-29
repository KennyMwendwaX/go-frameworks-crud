package db

import (
	"log"

	"github.com/kenny-mwendwa/go-restapi-crud/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DB_URI), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return db, nil
}
