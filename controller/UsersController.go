package controller

import (
	"database/sql"
	"encoding/json"
	"go-app/models"
	"log"
	"net/http"
	"strconv"
)

func GetUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := models.GetAllUsers(db)
		if err != nil {
			log.Println("Error getting users:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			log.Println("Error encoding response:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}


func CreateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Println("Error decoding request body:", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}


		err = models.CreateUser(db, &user)
		if err != nil {
			log.Println("Error creating user:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func UpdateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Println("Error parsing user ID:", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		var user models.User
		err = json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Println("Error decoding request body:", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

	
		existingUser, err := models.GetUserByID(db, id)
		if err != nil {
			log.Println("Error getting user:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		existingUser.Username = user.Username
		existingUser.Email = user.Email
		existingUser.Password = user.Password

		err = models.UpdateUser(db, &existingUser)
		if err != nil {
			log.Println("Error updating user:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Println("Error parsing user ID:", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		err = models.DeleteUser(db, id)
		if err != nil {
			log.Println("Error deleting user:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}