package routers

import (
	"net/http"

	"github.com/kenny-mwendwa/go-restapi-crud/handlers"
)

func SetUpBuiltInRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/users/create", handlers.CreateUser)
	mux.HandleFunc("/users", handlers.GetUsers)
	mux.HandleFunc("/users/:id", handlers.GetUser)

	// Add other routes if needed

	return mux
}
