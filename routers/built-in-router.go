package routers

import (
	"net/http"

	"github.com/kenny-mwendwa/go-restapi-crud/handlers"
)

func SetUpBuiltInRouter() *http.ServeMux {
	http.HandleFunc("/users", handlers.CreateUser)
	http.HandleFunc("/users", handlers.GetUsers)
	// Add other routes if needed
}
