package controller

import (
	"database/sql"
	"encoding/json"
	"go-app/handler"
	"log"
	"net/http"
	"strconv"
)

func GetUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := handler.GetAllUsers(db)
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
		var user handler.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Println("Error decoding request body:", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}


		err = handler.CreateUser(db, &user)
		if err != nil {
			log.Println("Error creating user:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			log.Println("Error encoding response:", err)
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

		var user handler.User
		err = json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Println("Error decoding request body:", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

	
		existingUser, err := handler.GetUserByID(db, id)
		if err != nil {
			log.Println("Error getting user:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		existingUser.Username = user.Username
		existingUser.Email = user.Email
		existingUser.Password = user.Password

		err = handler.UpdateUser(db, &existingUser)
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

		err = handler.DeleteUser(db, id)
		if err != nil {
			log.Println("Error deleting user:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}


func GetUserByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			log.Println("Missing user ID")
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Println("Error parsing user ID:", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		var user handler.User
		err = db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Println("User not found")
				http.NotFound(w, r)
				return
			}
			log.Println("Error getting user:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			log.Println("Error encoding response:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}