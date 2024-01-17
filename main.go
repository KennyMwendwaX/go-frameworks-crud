package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kenny-mwendwa/go-restapi-crud/routers"
)

func main() {
	// db, err := db.ConnectDB()

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// Migrate the schema
	// db.AutoMigrate(&models.User{})
	standardRouter := routers.StandardRouter()
	muxRouter := routers.MuxRouter()
	go func() {
		log.Fatal(http.ListenAndServe(":8000", standardRouter))
	}()
	fmt.Println("Server running at http://localhost:8001")

	go func() {
		log.Fatal(http.ListenAndServe(":8080", muxRouter))
	}()

	fmt.Println("Server running at http://localhost:8080")

	select {}
}
