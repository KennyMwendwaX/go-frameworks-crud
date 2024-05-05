package routers

import (
	"github.com/gorilla/mux"
	"github.com/kenny-mwendwa/go-crud-frameworks/internals/handlers"
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
