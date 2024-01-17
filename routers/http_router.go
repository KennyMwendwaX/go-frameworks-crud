package routers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/kenny-mwendwa/go-restapi-crud/handlers"
)

func HttpRouter() *httprouter.Router {
	r := httprouter.New()
	r.GET("/users", handlers.HttpGetUsers)
	r.POST("/users", handlers.HttpCreateUser)
	r.GET("/users/:id", handlers.HttpGetUser)
	return r
}
