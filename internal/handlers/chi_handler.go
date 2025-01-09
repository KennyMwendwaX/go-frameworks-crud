package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/KennyMwendwaX/go-frameworks-crud/internal/config"
	"github.com/KennyMwendwaX/go-frameworks-crud/internal/db"
	"github.com/KennyMwendwaX/go-frameworks-crud/internal/models"
	"github.com/KennyMwendwaX/go-frameworks-crud/internal/utils"
	"github.com/go-chi/chi/v5"
)

// CREATE USER
func ChiCreateUser(cfg *config.APIConfig) http.HandlerFunc {
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
		user, err := cfg.DB.CreateUser(r.Context(), db.CreateUserParams{
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
func ChiGetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer conn.Close(ctx)

	query := db.New(conn)

	users, err := query.GetUsers(ctx)
	if err != nil {
		log.Println("Error fetching users from the database:", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	usersJSON, err := json.Marshal(users)

	if err != nil {
		log.Println("Error marshaling users to JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write JSON response to the client
	w.WriteHeader(http.StatusOK)
	w.Write(usersJSON)
}

// GET ONE USER
func ChiGetUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer conn.Close(ctx)

	query := db.New(conn)

	// Extract user ID from request URL parameters
	userIdStr := chi.URLParam(r, "id")

	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		log.Println("Error converting userId to integer:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Query the DB for the user with the specified ID
	user, err := query.GetUser(ctx, int32(userId))
	if err != nil {
		log.Println("Error fetching user from the database:", err.Error())
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	userJSON, err := json.Marshal(user)

	if err != nil {
		log.Println("Error marshaling user to JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}

// UPDATE USER
func ChiUpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer conn.Close(ctx)

	query := db.New(conn)

	// Extract user ID from request URL parameters
	userIdStr := chi.URLParam(r, "id")

	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		log.Println("Error converting userId to integer:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	existingUser, err := query.GetUser(ctx, int32(userId))
	if err != nil {
		log.Println("Error fetching user from the database:", err.Error())
		http.Error(w, "User not found", http.StatusNotFound)
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
			log.Println("Error converting age to integer:", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
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
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DELETE USER
func ChiDeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer conn.Close(ctx)

	query := db.New(conn)

	// Extract user ID from request URL parameters
	userIdStr := chi.URLParam(r, "id")

	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		log.Println("Error converting userId to integer:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	_, err = query.GetUser(ctx, int32(userId))
	if err != nil {
		log.Println("Error fetching user from the database:", err.Error())
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err := query.DeleteUser(ctx, int32(userId)); err != nil {
		log.Println("Error deleting user from the database:", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
