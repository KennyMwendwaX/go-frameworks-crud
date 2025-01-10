package routers

import (
	"net/http"

	"github.com/KennyMwendwaX/go-frameworks-crud/internal/config"
	"github.com/KennyMwendwaX/go-frameworks-crud/internal/handlers"
)

func StandardRouter() *http.ServeMux {
	r := http.NewServeMux()
	cfg := config.ApiCfg()

	r.HandleFunc("GET /users", handlers.StandardGetUsers(cfg))
	r.HandleFunc("POST /users", handlers.StandardCreateUser(cfg))
	r.HandleFunc("GET /users/{id}", handlers.StandardGetUser(cfg))
	r.HandleFunc("PUT /users/{id}", handlers.StandardUpdateUser(cfg))
	r.HandleFunc("DELETE /users/{id}", handlers.StandardDeleteUser(cfg))

	return r
}
