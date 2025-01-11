package routers

import (
	"github.com/KennyMwendwaX/go-frameworks-crud/internal/config"
	"github.com/KennyMwendwaX/go-frameworks-crud/internal/handlers"
	"github.com/gin-gonic/gin"
)

func GinRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	cfg := config.ApiCfg()

	r.GET("/users", handlers.GinGetUsers(cfg))
	r.POST("/users", handlers.GinCreateUser(cfg))
	r.GET("/users/:id", handlers.GinGetUser(cfg))
	r.PUT("/users/:id", handlers.GinUpdateUser(cfg))
	r.DELETE("/users/:id", handlers.GinDeleteUser(cfg))

	return r
}
