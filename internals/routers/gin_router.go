package routers

import (
	"github.com/KennyMwendwaX/go-frameworks-crud/internals/handlers"
	"github.com/gin-gonic/gin"
)

func GinRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/users", handlers.GinGetUsers)
	r.POST("/users", handlers.GinCreateUser)
	r.GET("/users/:id", handlers.GinGetUser)
	r.PUT("/users/:id", handlers.GinUpdateUser)
	r.DELETE("/users/:id", handlers.GinDeleteUser)

	return r
}
