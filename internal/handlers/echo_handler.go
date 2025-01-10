package handlers

import (
	"net/http"
	"strconv"

	"github.com/KennyMwendwaX/go-frameworks-crud/internal/config"
	"github.com/KennyMwendwaX/go-frameworks-crud/internal/database"
	"github.com/labstack/echo/v4"
)

// CREATE USER
func EchoCreateUser(cfg *config.APIConfig) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		ageStr := c.FormValue("age")

		// Guard clauses to check if values are empty
		if name == "" || email == "" || ageStr == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request: Empty values"})
		}

		age, err := strconv.ParseUint(ageStr, 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request: Invalid age"})
		}

		user, err := cfg.DB.CreateUser(c.Request().Context(), database.CreateUserParams{
			Name:  name,
			Email: email,
			Age:   int32(age),
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error creating user"})
		}

		return c.JSON(http.StatusCreated, user)
	}
}

// GET ALL USERS
func EchoGetUsers(cfg *config.APIConfig) echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := cfg.DB.GetUsers(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		}

		return c.JSON(http.StatusOK, users)
	}
}

// GET ONE USER
func EchoGetUser(cfg *config.APIConfig) echo.HandlerFunc {
	return func(c echo.Context) error {
		userIdStr := c.Param("id")

		userId, err := strconv.ParseUint(userIdStr, 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request: Invalid user ID"})
		}

		user, err := cfg.DB.GetUser(c.Request().Context(), int32(userId))
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}

		return c.JSON(http.StatusOK, user)
	}
}

// UPDATE USER
func EchoUpdateUser(cfg *config.APIConfig) echo.HandlerFunc {
	return func(c echo.Context) error {
		userIdStr := c.Param("id")

		userId, err := strconv.ParseUint(userIdStr, 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request: Invalid user ID"})
		}

		existingUser, err := cfg.DB.GetUser(c.Request().Context(), int32(userId))
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
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
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request: Invalid age"})
			}
			existingUser.Age = int32(age)
		}

		updatedUser, err := cfg.DB.UpdateUser(c.Request().Context(), database.UpdateUserParams{
			ID:    existingUser.ID,
			Name:  existingUser.Name,
			Email: existingUser.Email,
			Age:   existingUser.Age,
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		}

		return c.JSON(http.StatusOK, updatedUser)
	}
}

// DELETE USER
func EchoDeleteUser(cfg *config.APIConfig) echo.HandlerFunc {
	return func(c echo.Context) error {
		userIdStr := c.Param("id")

		userId, err := strconv.ParseUint(userIdStr, 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request: Invalid user ID"})
		}

		_, err = cfg.DB.GetUser(c.Request().Context(), int32(userId))
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}

		if err := cfg.DB.DeleteUser(c.Request().Context(), int32(userId)); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		}

		return c.JSON(http.StatusNoContent, map[string]string{"message": "Successfully deleted the user"})
	}
}
