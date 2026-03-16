package handlers

import (

	  "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
    "ExpenseTracker/models"
)

type AuthHandler struct {
	DB *sql.DB
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{DB: db}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}

	if body.Username == "" || body.Email == "" || body.Password == "" {
		http.Error(w, `{"error":"username, email and password are required"}`, http.StatusBadRequest)
		return
	}

	if len(body.Password) < 6 {
		http.Error(w, `{"error":"password must be at least 6 characters"}`, http.StatusBadRequest)
		return
	}

	hashed, err :=bcrypt.GenerateFromPassword([]byte(body.Password),12)

	if err !=nil{
		log.Printf("Failed to create the password :%v",err)

		 http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
        return
	}

}
