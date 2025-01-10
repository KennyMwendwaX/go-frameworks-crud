package routers

import (
	"github.com/KennyMwendwaX/go-frameworks-crud/internal/config"
	"github.com/KennyMwendwaX/go-frameworks-crud/internal/handlers"
	"github.com/gorilla/mux"
)

func MuxRouter() *mux.Router {
	r := mux.NewRouter()
	cfg := config.ApiCfg()

	r.HandleFunc("/users", handlers.MuxGetUsers(cfg)).Methods("GET")
	r.HandleFunc("/users", handlers.MuxCreateUser(cfg)).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.MuxGetUser(cfg)).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.MuxUpdateUser(cfg)).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.MuxDeleteUser(cfg)).Methods("DELETE")

	return r
}
