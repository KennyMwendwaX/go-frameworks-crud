package routers

import (
	"github.com/KennyMwendwaX/go-frameworks-crud/internal/config"
	"github.com/KennyMwendwaX/go-frameworks-crud/internal/handlers"
	"github.com/labstack/echo/v4"
)

func EchoRouter() *echo.Echo {
	r := echo.New()
	cfg := config.ApiCfg()

	r.GET("/users", handlers.EchoGetUsers(cfg))
	r.POST("/users", handlers.EchoCreateUser(cfg))
	r.GET("/users/:id", handlers.EchoGetUser(cfg))
	r.PUT("/users/:id", handlers.EchoUpdateUser(cfg))
	r.DELETE("/users/:id", handlers.EchoDeleteUser(cfg))

	return r
}
