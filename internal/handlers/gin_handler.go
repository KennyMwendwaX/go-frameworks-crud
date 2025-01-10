package handlers

import (
	"net/http"
	"strconv"

	"github.com/KennyMwendwaX/go-frameworks-crud/internal/config"
	"github.com/KennyMwendwaX/go-frameworks-crud/internal/database"
	"github.com/KennyMwendwaX/go-frameworks-crud/internal/models"
	"github.com/gin-gonic/gin"
)

// CREATE USER
func GinCreateUser(cfg *config.APIConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse form values
		name := c.PostForm("name")
		email := c.PostForm("email")
		ageStr := c.PostForm("age")

		// Validate input
		if name == "" || email == "" || ageStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request: Empty values"})
			return
		}

		// Convert age to integer
		age, err := strconv.ParseUint(ageStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request: Invalid age format"})
			return
		}

		// Create user
		user, err := cfg.DB.CreateUser(c.Request.Context(), database.CreateUserParams{
			Name:  name,
			Email: email,
			Age:   int32(age),
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		c.JSON(http.StatusCreated, models.FromDatabaseUser(user))
	}
}

// GET ALL USERS
func GinGetUsers(cfg *config.APIConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := cfg.DB.GetUsers(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		c.JSON(http.StatusOK, models.FromDatabaseUsers(users))
	}
}

// GET ONE USER
func GinGetUser(cfg *config.APIConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract user ID from the path
		userIdStr := c.Param("id")

		userId, err := strconv.ParseUint(userIdStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request: Invalid user ID format"})
			return
		}

		// Get user by ID
		user, err := cfg.DB.GetUser(c.Request.Context(), int32(userId))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, models.FromDatabaseUser(user))
	}
}

// UPDATE USER
func GinUpdateUser(cfg *config.APIConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract user ID from the path
		userIdStr := c.Param("id")

		userId, err := strconv.ParseUint(userIdStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request: Invalid user ID format"})
			return
		}

		// Fetch the existing user
		existingUser, err := cfg.DB.GetUser(c.Request.Context(), int32(userId))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		// Parse form values
		name := c.PostForm("name")
		email := c.PostForm("email")
		ageStr := c.PostForm("age")

		// Update fields if provided
		if name != "" {
			existingUser.Name = name
		}
		if email != "" {
			existingUser.Email = email
		}
		if ageStr != "" {
			age, err := strconv.ParseUint(ageStr, 10, 32)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request: Invalid age format"})
				return
			}
			existingUser.Age = int32(age)
		}

		// Save updated user
		updatedUser, err := cfg.DB.UpdateUser(c.Request.Context(), database.UpdateUserParams{
			ID:    existingUser.ID,
			Name:  existingUser.Name,
			Email: existingUser.Email,
			Age:   existingUser.Age,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		c.JSON(http.StatusOK, models.FromDatabaseUser(updatedUser))
	}
}

// DELETE USER
func GinDeleteUser(cfg *config.APIConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract user ID from the path
		userIdStr := c.Param("id")

		userId, err := strconv.ParseUint(userIdStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request: Invalid user ID format"})
			return
		}

		// Check if user exists
		_, err = cfg.DB.GetUser(c.Request.Context(), int32(userId))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		// Delete the user
		if err := cfg.DB.DeleteUser(c.Request.Context(), int32(userId)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		c.JSON(http.StatusNoContent, gin.H{"message": "Successfully deleted the user"})
	}
}
