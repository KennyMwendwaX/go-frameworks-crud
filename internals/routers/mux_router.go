package routers

import (
	"github.com/KennyMwendwaX/go-frameworks-crud/internals/handlers"
	"github.com/gorilla/mux"
)

func MuxRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", handlers.MuxGetUsers).Methods("GET")
	r.HandleFunc("/users", handlers.MuxCreateUser).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.MuxGetUser).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.MuxUpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.MuxDeleteUser).Methods("DELETE")

	return r
}
