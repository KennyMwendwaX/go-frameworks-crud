package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/KennyMwendwaX/go-frameworks-crud/internals/db"
	"github.com/julienschmidt/httprouter"
)

// CREATE USER
func HttpCreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	// Convert age to integer
	age, err := strconv.ParseUint(ageStr, 10, 32)
	if err != nil {
		log.Println("Error converting age to integer:", err)
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
func HttpGetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := context.Background()

	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer conn.Close(ctx)

	query := db.New(conn)

	// Query the database for all users
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
func HttpGetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	userIdStr := ps.ByName("id")

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
func HttpUpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	userIdStr := ps.ByName("id")

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
func HttpDeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	userIdStr := ps.ByName("id")

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
