package routers

import (
	"github.com/kenny-mwendwa/go-restapi-crud/internal/handlers"
	"github.com/labstack/echo/v4"
)

func EchoRouter() *echo.Echo {
	e := echo.New()

	e.GET("/users", handlers.EchoGetUsers)
	e.POST("/users", handlers.EchoCreateUser)
	e.GET("/users/:id", handlers.EchoGetUser)
	e.PUT("/users/:id", handlers.EchoUpdateUser)
	e.DELETE("/users/:id", handlers.EchoDeleteUser)

	return e
}
