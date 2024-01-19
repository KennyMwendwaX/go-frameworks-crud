package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kenny-mwendwa/go-restapi-crud/db"
	"github.com/kenny-mwendwa/go-restapi-crud/models"
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
