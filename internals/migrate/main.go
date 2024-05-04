package main

import (
	"log"

	"github.com/kenny-mwendwa/go-restapi-crud/internal/db"
	"github.com/kenny-mwendwa/go-restapi-crud/internal/models"
)

func main() {
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&models.User{})
}
