package users

import (
	"ExpenseTracker/config"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(username, email, hashedPassword string) error {
	_, err := config.DB.Exec(
		"INSERT INTO users(username, email, password) VALUES ( $1, $2, $3)", username, email, hashedPassword)
	return err
}

func GetUserByEmail(email string) (User, error) {
	var user User
	row := config.DB.QueryRow(
		"SELECT id, username ,email, password FROM users WHERE email=$1", email)
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	return user, err
}
