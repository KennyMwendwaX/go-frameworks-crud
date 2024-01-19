package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kenny-mwendwa/go-restapi-crud/handlers"
)

func GinRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/users", handlers.GinGetUsers)
	r.POST("/users", handlers.GinCreateUser)
	r.GET("/users/:id", handlers.GinGetUser)
	r.PUT("/users/:id", handlers.GinUpdateUser)

	return r
}
