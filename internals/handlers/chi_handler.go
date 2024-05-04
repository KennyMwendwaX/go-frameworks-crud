package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kenny-mwendwa/go-restapi-crud/internals/db"
	"github.com/kenny-mwendwa/go-restapi-crud/internals/models"
)

// CREATE USER
func ChiCreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer conn.Close(ctx)

	query := db.New(conn)

	// Parse form data
	err = r.ParseForm()
	if err != nil {
		log.Println("Error parsing form data:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Get form values
	name := r.FormValue("name")
	email := r.FormValue("email")
	ageStr := r.FormValue("age")

	// Guard clauses to check if values are empty
	if name == "" || email == "" || ageStr == "" {
		log.Println("Empty values detected")
		http.Error(w, "Bad Request: Empty values", http.StatusBadRequest)
		return
	}

	// Convert age to unit
	age, err := strconv.ParseUint(ageStr, 10, 32)
	if err != nil {
		log.Println("Error converting age to unit:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Create user
	if err := query.CreateUser(ctx, db.CreateUserParams{
		Name:  name,
		Email: email,
		Age:   int32(age),
	}); err != nil {
		log.Println("Error creating user:", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusCreated)
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

	userId, err := strconv.ParseInt(userIdStr, 10, 32)
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
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Extract user ID from request URL parameters
	userIdStr := chi.URLParam(r, "id")

	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		log.Println("Error converting userId to unit32:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var existingUser models.User
	result := db.First(&existingUser, userId)
	if result.Error != nil {
		log.Println("Error fetching user from the database:", result.Error)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Get form values
	name := r.FormValue("name")
	email := r.FormValue("email")
	ageStr := r.FormValue("age")

	age, err := strconv.ParseUint(ageStr, 10, 32)
	if err != nil {
		log.Println("Error converting age to unit32:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
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

	w.WriteHeader(http.StatusOK)
}

// DELETE USER
func ChiDeleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Extract user ID from request URL parameters
	userIdStr := chi.URLParam(r, "id")

	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		log.Println("Error converting userId to unit32:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var existingUser models.User
	result := db.First(&existingUser, userId)
	if result.Error != nil {
		log.Println("Error fetching user from the database:", result.Error)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	db.Delete(&existingUser)

	w.WriteHeader(http.StatusNoContent)
}
