package routers

import (
	"net/http"

	"github.com/kenny-mwendwa/go-restapi-crud/handlers"
)

func StandardRouter() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/users", handlers.StandardGetUsers)
	r.HandleFunc("/users/create", handlers.StandardCreateUser)
	r.HandleFunc("/users/", handlers.StandardGetUser)
	r.HandleFunc("/users/update/", handlers.StandardUpdateUser)
	r.HandleFunc("/users/delete/", handlers.StandardDeleteUser)

	return r
}
