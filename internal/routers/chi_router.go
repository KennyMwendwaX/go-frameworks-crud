package routers

import (
	"github.com/KennyMwendwaX/go-frameworks-crud/internal/config"
	"github.com/KennyMwendwaX/go-frameworks-crud/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func ChiRouter() *chi.Mux {
	cfg := config.ApiCfg()
	r := chi.NewRouter()

	r.Get("/users", handlers.ChiGetUsers(cfg))
	r.Post("/users", handlers.ChiCreateUser(cfg))
	r.Get("/users/{id}", handlers.ChiGetUser(cfg))
	r.Put("/users/{id}", handlers.ChiUpdateUser(cfg))
	r.Delete("/users/{id}", handlers.ChiDeleteUser(cfg))

	return r
}
