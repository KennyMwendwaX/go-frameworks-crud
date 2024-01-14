package routers

import (
	"github.com/gorilla/mux"
	"github.com/kenny-mwendwa/go-restapi-crud/handlers"
)

func SetUpBuiltInRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users/create", handlers.CreateUser)
	r.HandleFunc("/users", handlers.GetUsers)
	r.HandleFunc("/users/{id:[0-9]+}", handlers.GetUser)

	// Add other routes if needed

	return r
}
