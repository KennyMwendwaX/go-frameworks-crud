package routers

import (
	"github.com/KennyMwendwaX/go-frameworks-crud/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func ChiRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/users", handlers.ChiGetUsers)
	r.Post("/users", handlers.ChiCreateUser)
	r.Get("/users/{id}", handlers.ChiGetUser)
	r.Put("/users/{id}", handlers.ChiUpdateUser)
	r.Delete("/users/{id}", handlers.ChiDeleteUser)

	return r
}
