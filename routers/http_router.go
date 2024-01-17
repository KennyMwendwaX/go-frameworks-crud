package routers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/kenny-mwendwa/go-restapi-crud/handlers"
)

func HttpRouter() *httprouter.Router {
	r := httprouter.New()
	r.GET("/users", handlers.HttpGetUsers)
	return r
}
