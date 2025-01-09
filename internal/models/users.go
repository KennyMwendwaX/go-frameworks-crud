package models

import (
	"github.com/KennyMwendwaX/go-frameworks-crud/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID        int32              `json:"id"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Age       int32              `json:"age"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

func FromDatabaseUser(databaseUser database.User) User {
	return User{
		ID:        databaseUser.ID,
		Name:      databaseUser.Name,
		Email:     databaseUser.Email,
		Age:       databaseUser.Age,
		CreatedAt: databaseUser.CreatedAt,
	}
}
