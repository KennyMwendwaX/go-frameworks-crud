package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kenny-mwendwa/go-restapi-crud/internal/db"
	"github.com/kenny-mwendwa/go-restapi-crud/internal/models"
)

// CREATE USER
func GinCreateUser(c *gin.Context) {
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Retrieve data from the form
	name := c.PostForm("name")
	email := c.PostForm("email")
	ageStr := c.PostForm("age")

	// Validate input
	if name == "" || email == "" || ageStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request: Empty values"})
		return
	}

	// Convert age to uint
	age, err := strconv.ParseUint(ageStr, 10, 32)
	if err != nil {
		log.Println("Error converting age to uint:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request: Invalid age format"})
		return
	}

	// Create a new user
	newUser := models.User{
		Name:  name,
		Email: email,
		Age:   uint(age),
	}

	// Save the new user to the database
	result := db.Create(&newUser)
	if result.Error != nil {
		log.Println("Error creating user:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Return a success response
	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

// GET ALL USERS
func GinGetUsers(c *gin.Context) {
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var users []models.User

	result := db.Find(&users)
	if result.Error != nil {
		log.Println("Error fetching users from the database:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GET ONE USER
func GinGetUser(c *gin.Context) {
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Retrieve user ID from the URL parameters
	userIDStr := c.Param("id")

	// Validate user ID
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		log.Println("Error converting userID to uint:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request: Invalid user ID format"})
		return
	}

	// Query the database for the user with the specified ID
	var user models.User
	result := db.First(&user, userID)
	if result.Error != nil {
		log.Println("Error fetching user from the database:", result.Error)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Return JSON response to the client
	c.JSON(http.StatusOK, user)
}

// UPDATE USER
func GinUpdateUser(c *gin.Context) {
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Retrieve user ID from the URL parameters
	userIDStr := c.Param("id")

	// Validate user ID
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		log.Println("Error converting userID to uint:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request: Invalid user ID format"})
		return
	}

	var existingUser models.User

	result := db.First(&existingUser, userID)
	if result.Error != nil {
		log.Println("Error fetching user from the database:", result.Error)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Retrieve data from the form
	name := c.PostForm("name")
	email := c.PostForm("email")
	ageStr := c.PostForm("age")

	// Convert age to uint
	age, err := strconv.ParseUint(ageStr, 10, 32)
	if err != nil {
		log.Println("Error converting age to uint:", err)
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
		existingUser.Age = uint(age)
	}

	// Save the updated user to the database
	db.Save(&existingUser)

	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

// DELETE USER
func GinDeleteUser(c *gin.Context) {
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Retrieve user ID from the URL parameters
	userIDStr := c.Param("id")

	// Validate user ID
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		log.Println("Error converting userID to uint:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request: Invalid user ID format"})
		return
	}

	var existingUser models.User

	result := db.First(&existingUser, userID)
	if result.Error != nil {
		log.Println("Error fetching user from the database:", result.Error)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	db.Delete(&existingUser)

	c.JSON(http.StatusNoContent, gin.H{"message": "User deleted"})
}
