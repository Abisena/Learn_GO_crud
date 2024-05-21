package handler

import (
	"database/sql"
	"errors"
)

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}


func GetAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT ID, Username, Email, Password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func CreateUser(db *sql.DB, user *User) error {
	_, err := db.Exec("INSERT INTO users (Username, Email, Password, ID) VALUES ($1, $2, $3, $4)", user.Username, user.Email, user.Password, user.ID)
	return err
}

func UpdateUser(db *sql.DB, user *User) error {
	_, err := db.Exec("UPDATE users SET Username = $1, Email = $2, Password = $3 WHERE id = $4", user.Username, user.Email, user.Password, user.ID)
	return err
}

func DeleteUser(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func GetUserByID(db *sql.DB, id int) (User, error) {
	var user User
	err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}