package db

import (
	"fmt"
	"log"

	"github.com/kenny-mwendwa/go-restapi-crud/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.DBConfig.Host, config.DBConfig.User, config.DBConfig.Password, config.DBConfig.Name, config.DBConfig.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	// if err := db.Ping(); err != nil {
	// 	log.Println("Failed to ping database:", err)
	// 	return nil, err
	// }

	return db, nil
}
