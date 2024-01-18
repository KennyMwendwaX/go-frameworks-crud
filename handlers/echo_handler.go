package handlers

import (
	"log"
	"net/http"

	"github.com/kenny-mwendwa/go-restapi-crud/db"
	"github.com/kenny-mwendwa/go-restapi-crud/models"
	"github.com/labstack/echo/v4"
)

// GET ALL USERS
func EchoGetUsers(c echo.Context) error {
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	var users []models.User

	result := db.Find(&users)
	if result.Error != nil {
		log.Println("Error fetching users from the database:", result.Error)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	// Set Content-Type header
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	return c.JSON(http.StatusOK, users)
}
