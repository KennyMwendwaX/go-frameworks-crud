package routers

import (
	"github.com/gorilla/mux"
	"github.com/kenny-mwendwa/go-restapi-crud/handlers"
)

func SetUpBuiltInRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")

	// Add other routes if needed

	return r
}
