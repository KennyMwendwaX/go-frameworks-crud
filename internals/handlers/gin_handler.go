package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kenny-mwendwa/go-restapi-crud/internals/db"
)

// CREATE USER
func GinCreateUser(c *gin.Context) {
	ctx := context.Background()

	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer conn.Close(ctx)

	query := db.New(conn)

	// Retrieve data from the form
	name := c.PostForm("name")
	email := c.PostForm("email")
	ageStr := c.PostForm("age")

	// Validate input
	if name == "" || email == "" || ageStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request: Empty values"})
		return
	}

	// Convert age to integer
	age, err := strconv.ParseInt(ageStr, 10, 32)
	if err != nil {
		log.Println("Error converting age to integer:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request: Invalid age format"})
		return
	}

	// Create user
	if err := query.CreateUser(ctx, db.CreateUserParams{
		Name:  name,
		Email: email,
		Age:   int32(age),
	}); err != nil {
		log.Println("Error creating user:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Return a success response
	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

// GET ALL USERS
func GinGetUsers(c *gin.Context) {
	ctx := context.Background()

	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer conn.Close(ctx)

	query := db.New(conn)

	users, err := query.GetUsers(ctx)
	if err != nil {
		log.Println("Error fetching users from the database:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GET ONE USER
func GinGetUser(c *gin.Context) {
	ctx := context.Background()

	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer conn.Close(ctx)

	query := db.New(conn)

	// Retrieve user ID from the URL parameters
	userIdStr := c.Param("id")

	// Validate user ID
	userId, err := strconv.ParseInt(userIdStr, 10, 32)
	if err != nil {
		log.Println("Error converting userId to integer:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request: Invalid user ID format"})
		return
	}

	// Query the database for the user with the specified ID
	user, err := query.GetUser(ctx, int32(userId))
	if err != nil {
		log.Println("Error fetching user from the database:", err.Error())
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Return JSON response to the client
	c.JSON(http.StatusOK, user)
}

// UPDATE USER
func GinUpdateUser(c *gin.Context) {
	ctx := context.Background()

	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer conn.Close(ctx)

	query := db.New(conn)

	// Retrieve user ID from the URL parameters
	userIdStr := c.Param("id")

	// Validate user ID
	userId, err := strconv.ParseInt(userIdStr, 10, 32)
	if err != nil {
		log.Println("Error converting userId to integer:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request: Invalid user ID format"})
		return
	}

	existingUser, err := query.GetUser(ctx, int32(userId))
	if err != nil {
		log.Println("Error fetching user from the database:", err.Error())
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Retrieve data from the form
	name := c.PostForm("name")
	email := c.PostForm("email")
	ageStr := c.PostForm("age")

	// Convert age to integer
	age, err := strconv.ParseInt(ageStr, 10, 32)
	if err != nil {
		log.Println("Error converting age to integer:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request: Invalid age format"})
		return
	}

	// Update user fields if provided
	if name != "" {
		existingUser.Name = name
	}

	if email != "" {
		existingUser.Email = email
	}

	if ageStr != "" {
		existingUser.Age = int32(age)
	}

	// Save the updated user to the database
	if err := query.UpdateUser(ctx, db.UpdateUserParams{
		Name:  name,
		Email: email,
		Age:   int32(age),
	}); err != nil {
		log.Println("Error updating user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

// DELETE USER
func GinDeleteUser(c *gin.Context) {
	ctx := context.Background()

	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer conn.Close(ctx)

	query := db.New(conn)

	// Retrieve user ID from the URL parameters
	userIdStr := c.Param("id")

	// Validate user ID
	userId, err := strconv.ParseInt(userIdStr, 10, 32)
	if err != nil {
		log.Println("Error converting userId to integer:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request: Invalid user ID format"})
		return
	}

	_, err = query.GetUser(ctx, int32(userId))
	if err != nil {
		log.Println("Error fetching user from the database:", err.Error)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := query.DeleteUser(ctx, int32(userId)); err != nil {
		log.Println("Error deleting user from the database:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "User deleted"})
}
