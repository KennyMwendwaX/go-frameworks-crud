package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kenny-mwendwa/go-restapi-crud/db"
	"github.com/kenny-mwendwa/go-restapi-crud/models"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

// READ
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var users []models.User

	db.Find(&users)

	usersJSON, err := json.Marshal(users)

	if err != nil {
		log.Println("Error marshaling users to JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write JSON response to the client
	w.Write(usersJSON)
}

// Add other handlers if needed
