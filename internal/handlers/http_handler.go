package handlers

import (
	"net/http"
	"strconv"

	"github.com/KennyMwendwaX/go-frameworks-crud/internal/config"
	"github.com/KennyMwendwaX/go-frameworks-crud/internal/database"
	"github.com/KennyMwendwaX/go-frameworks-crud/internal/utils"
	"github.com/julienschmidt/httprouter"
)

// CREATE USER
func HttpCreateUser(cfg *config.APIConfig) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		err := r.ParseForm()
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Bad Request: Invalid form data")
			return
		}

		name := r.FormValue("name")
		email := r.FormValue("email")
		ageStr := r.FormValue("age")

		if name == "" || email == "" || ageStr == "" {
			utils.RespondWithError(w, http.StatusBadRequest, "Bad Request: Empty values")
			return
		}

		age, err := strconv.ParseUint(ageStr, 10, 32)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Bad Request: Invalid age")
			return
		}

		user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
			Name:  name,
			Email: email,
			Age:   int32(age),
		})
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Error creating user")
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, user)
	}
}

// GET ALL USERS
func HttpGetUsers(cfg *config.APIConfig) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		users, err := cfg.DB.GetUsers(r.Context())
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, users)
	}
}

// GET ONE USER
func HttpGetUser(cfg *config.APIConfig) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		userIdStr := ps.ByName("id")
		userId, err := strconv.ParseUint(userIdStr, 10, 32)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Bad Request: Invalid user ID")
			return
		}

		user, err := cfg.DB.GetUser(r.Context(), int32(userId))
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "User not found")
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, user)
	}
}

// UPDATE USER
func HttpUpdateUser(cfg *config.APIConfig) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		userIdStr := ps.ByName("id")
		userId, err := strconv.ParseUint(userIdStr, 10, 32)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Bad Request: Invalid user ID")
			return
		}

		existingUser, err := cfg.DB.GetUser(r.Context(), int32(userId))
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "User not found")
			return
		}

		name := r.FormValue("name")
		email := r.FormValue("email")
		ageStr := r.FormValue("age")

		if name != "" {
			existingUser.Name = name
		}
		if email != "" {
			existingUser.Email = email
		}
		if ageStr != "" {
			age, err := strconv.ParseUint(ageStr, 10, 32)
			if err != nil {
				utils.RespondWithError(w, http.StatusBadRequest, "Bad Request: Invalid age")
				return
			}
			existingUser.Age = int32(age)
		}

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

		utils.RespondWithJSON(w, http.StatusOK, updatedUser)
	}
}

// DELETE USER
func HttpDeleteUser(cfg *config.APIConfig) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		userIdStr := ps.ByName("id")
		userId, err := strconv.ParseUint(userIdStr, 10, 32)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Bad Request: Invalid user ID")
			return
		}

		_, err = cfg.DB.GetUser(r.Context(), int32(userId))
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "User not found")
			return
		}

		err = cfg.DB.DeleteUser(r.Context(), int32(userId))
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		utils.RespondWithJSON(w, http.StatusNoContent, "Successfully deleted the user")
	}
}
