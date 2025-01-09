package handlers

import (
	"net/http"
	"strconv"

	"github.com/KennyMwendwaX/go-frameworks-crud/internal/config"
	"github.com/KennyMwendwaX/go-frameworks-crud/internal/database"
	"github.com/KennyMwendwaX/go-frameworks-crud/internal/models"
	"github.com/KennyMwendwaX/go-frameworks-crud/internal/utils"
	"github.com/gorilla/mux"
)

// CREATE USER
func MuxCreateUser(cfg *config.APIConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse form data
		err := r.ParseForm()
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		// Get form values
		name := r.FormValue("name")
		email := r.FormValue("email")
		ageStr := r.FormValue("age")

		// Guard clauses to check if values are empty
		if name == "" || email == "" || ageStr == "" {
			utils.RespondWithError(w, http.StatusBadRequest, "Bad Request: Empty values")
			return
		}

		// Convert age to integer
		age, err := strconv.ParseUint(ageStr, 10, 32)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		// Create user
		user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
			Name:  name,
			Email: email,
			Age:   int32(age),
		})
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Error creating user")
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, models.FromDatabaseUser(user))
	}
}

// GET ALL USERS
func MuxGetUsers(cfg *config.APIConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		users, err := cfg.DB.GetUsers(r.Context())
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, models.FromDatabaseUsers(users))
	}
}

// GET ONE USER
func MuxGetUser(cfg *config.APIConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract user ID from request URL parameters
		vars := mux.Vars(r)
		userIdStr, ok := vars["id"]
		if !ok {
			utils.RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		userId, err := strconv.ParseUint(userIdStr, 10, 32)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		// Query the DB for the user with the specified ID
		user, err := cfg.DB.GetUser(r.Context(), int32(userId))
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, models.FromDatabaseUser(user))
	}
}

// UPDATE USER
func MuxUpdateUser(cfg *config.APIConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Extract user ID from request URL parameters
		vars := mux.Vars(r)
		userIdStr, ok := vars["id"]
		if !ok {
			utils.RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		userId, err := strconv.ParseUint(userIdStr, 10, 32)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		existingUser, err := cfg.DB.GetUser(r.Context(), int32(userId))
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "User not found")
			return
		}

		// Get form values
		name := r.FormValue("name")
		email := r.FormValue("email")
		ageStr := r.FormValue("age")

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
				utils.RespondWithError(w, http.StatusBadRequest, "Bad Request")
				return
			}
			existingUser.Age = int32(age)
		}

		// Save the updated user to the database
		updatedUser, err := cfg.DB.UpdateUser(r.Context(), database.UpdateUserParams{
			ID:    existingUser.ID,
			Name:  existingUser.Name,
			Email: existingUser.Email,
			Age:   existingUser.Age,
		})
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, models.FromDatabaseUser(updatedUser))
	}
}

// DELETE USER
func MuxDeleteUser(cfg *config.APIConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract user ID from request URL parameters
		vars := mux.Vars(r)
		userIdStr, ok := vars["id"]
		if !ok {
			utils.RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		userId, err := strconv.ParseUint(userIdStr, 10, 32)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		_, err = cfg.DB.GetUser(r.Context(), int32(userId))
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "User not found")
			return
		}

		if err := cfg.DB.DeleteUser(r.Context(), int32(userId)); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		utils.RespondWithJSON(w, http.StatusNoContent, "Successfully deleted the user")
	}
}
