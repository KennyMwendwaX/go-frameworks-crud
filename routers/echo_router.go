package routers

import "github.com/labstack/echo/v4"

func EchoRouter() {
	e := echo.New()

	e.GET("/users")
}
