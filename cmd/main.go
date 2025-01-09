package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KennyMwendwaX/go-frameworks-crud/internal/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	standardRouter := routers.StandardRouter()
	httpRouter := routers.HttpRouter()
	muxRouter := routers.MuxRouter()
	chiRouter := routers.ChiRouter()
	e := routers.EchoRouter()
	r := routers.GinRouter()

	gin.SetMode(gin.ReleaseMode)

	// Standard lib routine
	go func() {
		log.Fatal(http.ListenAndServe(":8000", standardRouter))
	}()
	fmt.Println("Standard router running at http://localhost:8000")

	// httprouter routine
	go func() {
		log.Fatal(http.ListenAndServe(":8001", httpRouter))
	}()
	fmt.Println("HttpRouter running at http://localhost:8001")

	// Mux router routine
	go func() {
		log.Fatal(http.ListenAndServe(":8002", muxRouter))
	}()
	fmt.Println("Mux router running at http://localhost:8002")

	// Chi router routine
	go func() {
		log.Fatal(http.ListenAndServe(":8003", chiRouter))
	}()
	fmt.Println("Chi router running at http://localhost:8003")

	// Echo routine
	go func() {
		if err := e.Start(":8004"); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	fmt.Println("Echo running at http://localhost:8004")

	// Gin routine
	go func() {
		r.Run(":8005")
	}()
	fmt.Println("Gin running at http://localhost:8005")

	select {}
}
