package routers

import (
	"github.com/KennyMwendwaX/go-frameworks-crud/internals/handlers"
	"github.com/julienschmidt/httprouter"
)

func HttpRouter() *httprouter.Router {
	r := httprouter.New()

	r.GET("/users", handlers.HttpGetUsers)
	r.POST("/users", handlers.HttpCreateUser)
	r.GET("/users/:id", handlers.HttpGetUser)
	r.PUT("/users/:id", handlers.HttpUpdateUser)
	r.DELETE("/users/:id", handlers.HttpDeleteUser)

	return r
}
