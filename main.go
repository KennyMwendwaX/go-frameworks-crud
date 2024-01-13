package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kenny-mwendwa/go-restapi-crud/db"
	"github.com/kenny-mwendwa/go-restapi-crud/models"
	"github.com/kenny-mwendwa/go-restapi-crud/routers"
)

func main() {
	db, err := db.ConnectDB()

	if err != nil {
		log.Fatal(err.Error())
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})
	builtInRouter := routers.SetUpBuiltInRouter()
	go func() {
		log.Fatal(http.ListenAndServe(":8080", builtInRouter))
	}()

	fmt.Println("Server running at http://localhost:8080")

	select {}
}
