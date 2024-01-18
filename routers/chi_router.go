package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/kenny-mwendwa/go-restapi-crud/handlers"
)

func ChiRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/users", handlers.ChiGetUsers)
	r.Post("/users", handlers.ChiCreateUser)
	r.Get("/users/{id}", handlers.ChiGetUser)
	r.Put("/users/{id}", handlers.ChiUpdateUser)

	return r
}
