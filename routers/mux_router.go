package routers

import (
	"github.com/gorilla/mux"
	"github.com/kenny-mwendwa/go-restapi-crud/handlers"
)

func SetUpBuiltInRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", handlers.CreateUser)
	r.HandleFunc("/users", handlers.GetUsers)
	r.HandleFunc("/users/{id:[0-9]+}", handlers.GetUser)
	r.HandleFunc("/users/{id:[0-9]+}", handlers.UpdateUser)

	// Add other routes if needed

	return r
}
