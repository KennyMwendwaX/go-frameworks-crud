package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var DB_URI string

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set configuration variables
	host := os.Getenv("HOST")
	portStr := os.Getenv("PORT")
	port, err := strconv.ParseInt(portStr, 10, 64)
	if err != nil {
		log.Fatal("Error converting PORT to integer:", err)
	}
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	name := os.Getenv("NAME")
	DB_URI = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, name, port)
}
