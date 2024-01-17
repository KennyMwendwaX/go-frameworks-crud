package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kenny-mwendwa/go-restapi-crud/db"
	"github.com/kenny-mwendwa/go-restapi-crud/models"
)

// GET ALL USERS
func ChiGetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var users []models.User

	// Query the database for all users
	result := db.Find(&users)
	if result.Error != nil {
		log.Println("Error fetching users from the database:", result.Error)
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
