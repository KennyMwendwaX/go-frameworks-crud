package db

import (
	"log"

	"github.com/kenny-mwendwa/go-restapi-crud/models"
)

func MigrateDB() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&models.User{})
}
