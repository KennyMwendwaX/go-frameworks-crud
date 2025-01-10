package routers

import (
	"github.com/KennyMwendwaX/go-frameworks-crud/internal/config"
	"github.com/KennyMwendwaX/go-frameworks-crud/internal/handlers"
	"github.com/julienschmidt/httprouter"
)

func HttpRouter() *httprouter.Router {
	r := httprouter.New()
	cfg := config.ApiCfg()

	r.GET("/users", handlers.HttpGetUsers(cfg))
	r.POST("/users", handlers.HttpCreateUser(cfg))
	r.GET("/users/:id", handlers.HttpGetUser(cfg))
	r.PUT("/users/:id", handlers.HttpUpdateUser(cfg))
	r.DELETE("/users/:id", handlers.HttpDeleteUser(cfg))

	return r
}
