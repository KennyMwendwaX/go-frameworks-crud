package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kenny-mwendwa/go-restapi-crud/routers"
)

func main() {
	routers.SetUpBuiltInRouter()
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	fmt.Println("Server running at http://localhost:8080")

	select {}
}
