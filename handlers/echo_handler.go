package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/kenny-mwendwa/go-restapi-crud/db"
	"github.com/kenny-mwendwa/go-restapi-crud/models"
	"github.com/labstack/echo/v4"
)

// CREATE USER
func EchoCreateUser(c echo.Context) error {
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	name := c.FormValue("name")
	email := c.FormValue("email")
	ageStr := c.FormValue("age")

	// Guard clauses to check if values are empty
	if name == "" || email == "" || ageStr == "" {
		log.Println("Empty values detected")
		return c.String(http.StatusBadRequest, "Bad Request: Empty values")
	}

	age, err := strconv.ParseUint(ageStr, 10, 32)
	if err != nil {
		log.Println("Error converting age to unit:", err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	newUser := models.User{
		Name:  name,
		Email: email,
		Age:   uint(age),
	}

	// Create user
	result := db.Create(&newUser)
	if result.Error != nil {
		log.Println("Error creating user:", result.Error)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	// Return a success response
	return c.String(http.StatusCreated, "User created")
}

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
