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
	httpRouter := routers.HttpRouter()
	muxRouter := routers.MuxRouter()

	go func() {
		log.Fatal(http.ListenAndServe(":8000", standardRouter))
	}()
	fmt.Println("Standard router running at http://localhost:8000")

	go func() {
		log.Fatal(http.ListenAndServe(":8001", httpRouter))
	}()
	fmt.Println("HttpRouter running at http://localhost:8001")

	go func() {
		log.Fatal(http.ListenAndServe(":8002", muxRouter))
	}()
	fmt.Println("Mux router running at http://localhost:8002")

	select {}
}
