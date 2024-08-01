package routers

import (
	"github.com/KennyMwendwaX/go-frameworks-crud/internals/handlers"
	"github.com/labstack/echo/v4"
)

func EchoRouter() *echo.Echo {
	r := echo.New()

	r.GET("/users", handlers.EchoGetUsers)
	r.POST("/users", handlers.EchoCreateUser)
	r.GET("/users/:id", handlers.EchoGetUser)
	r.PUT("/users/:id", handlers.EchoUpdateUser)
	r.DELETE("/users/:id", handlers.EchoDeleteUser)

	return r
}
