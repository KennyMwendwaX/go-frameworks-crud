package routers

import (
	"github.com/kenny-mwendwa/go-restapi-crud/handlers"
	"github.com/labstack/echo/v4"
)

func EchoRouter() *echo.Echo {
	e := echo.New()

	e.GET("/users", handlers.EchoGetUsers)

	return e
}
