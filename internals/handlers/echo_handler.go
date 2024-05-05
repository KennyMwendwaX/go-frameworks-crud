package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/kenny-mwendwa/go-crud-frameworks/internals/db"
	"github.com/labstack/echo/v4"
)

// CREATE USER
func EchoCreateUser(c echo.Context) error {
	ctx := context.Background()

	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	defer conn.Close(ctx)

	query := db.New(conn)

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

	// Create user
	if err := query.CreateUser(ctx, db.CreateUserParams{
		Name:  name,
		Email: email,
		Age:   int32(age),
	}); err != nil {
		log.Println("Error creating user:", err.Error())
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	// Return a success response
	return c.String(http.StatusCreated, "User created")
}

// GET ALL USERS
func EchoGetUsers(c echo.Context) error {
	ctx := context.Background()

	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	defer conn.Close(ctx)

	query := db.New(conn)

	users, err := query.GetUsers(ctx)
	if err != nil {
		log.Println("Error fetching users from the database:", err.Error())
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	// Set Content-Type header
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	return c.JSON(http.StatusOK, users)
}

// GET ONE USER
func EchoGetUser(c echo.Context) error {
	ctx := context.Background()

	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	defer conn.Close(ctx)

	query := db.New(conn)

	// Extract user ID from request URL parameters
	userIdStr := c.Param("id")

	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		log.Println("Error converting userId to integer:", err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	// Query the DB for the user with the specified ID
	user, err := query.GetUser(ctx, int32(userId))
	if err != nil {
		log.Println("Error fetching user from the database:", err.Error())
		return c.String(http.StatusNotFound, "User not found")
	}

	// Set Content-Type header
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	return c.JSON(http.StatusOK, user)
}

// UPDATE USER
func EchoUpdateUser(c echo.Context) error {
	ctx := context.Background()

	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	defer conn.Close(ctx)

	query := db.New(conn)

	// Extract user ID from request URL parameters
	userIdStr := c.Param("id")

	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		log.Println("Error converting userId to integer:", err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	existingUser, err := query.GetUser(ctx, int32(userId))
	if err != nil {
		log.Println("Error fetching user from the database:", err.Error())
		return c.String(http.StatusNotFound, "User not found")
	}

	name := c.FormValue("name")
	email := c.FormValue("email")
	ageStr := c.FormValue("age")

	// Update user fields if provided
	if name != "" {
		existingUser.Name = name
	}
	if email != "" {
		existingUser.Email = email
	}
	if ageStr != "" {
		age, err := strconv.ParseUint(ageStr, 10, 32)
		if err != nil {
			log.Println("Error converting age to integer:", err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		existingUser.Age = int32(age)
	}

	// Save the updated user to the database
	if err := query.UpdateUser(ctx, db.UpdateUserParams{
		ID:    existingUser.ID,
		Name:  existingUser.Name,
		Email: existingUser.Email,
		Age:   existingUser.Age,
	}); err != nil {
		log.Println("Error updating user:", err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.String(http.StatusOK, "User updated")
}

// DELETE USER
func EchoDeleteUser(c echo.Context) error {
	ctx := context.Background()

	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	defer conn.Close(ctx)

	query := db.New(conn)

	// Extract user ID from request URL parameters
	userIdStr := c.Param("id")

	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		log.Println("Error converting userId to integer:", err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	_, err = query.GetUser(ctx, int32(userId))
	if err != nil {
		log.Println("Error fetching user from the database:", err.Error())
		return c.String(http.StatusNotFound, "User not found")
	}

	if err := query.DeleteUser(ctx, int32(userId)); err != nil {
		log.Println("Error deleting user from the database:", err.Error())
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.String(http.StatusNoContent, "User deleted")
}
