package main

import (
	"log"

	"github.com/kenny-mwendwa/go-restapi-crud/db"
	"github.com/kenny-mwendwa/go-restapi-crud/models"
)

func main() {
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&models.User{})
}
