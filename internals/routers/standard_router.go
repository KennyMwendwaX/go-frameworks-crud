package routers

import (
	"net/http"

	"github.com/kenny-mwendwa/go-crud-frameworks/internals/handlers"
)

func StandardRouter() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("GET /users", handlers.StandardGetUsers)
	r.HandleFunc("POST /users", handlers.StandardCreateUser)
	r.HandleFunc("GET /users/{id}", handlers.StandardGetUser)
	r.HandleFunc("PUT /users/{id}", handlers.StandardUpdateUser)
	r.HandleFunc("DELETE /users/{id}", handlers.StandardDeleteUser)

	return r
}
