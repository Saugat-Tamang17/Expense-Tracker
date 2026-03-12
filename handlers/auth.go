package handlers

import (
	"ExpenseTracker/models"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	hashed, err := bcrypt.GenerateFromPassword([]byte(body.Password), 12)
	if err != nil {
		http.Error(w, "Error Hashing Password", http.StatusInternalServerError)
		return

	}

	err = models.CreateUser(body.Username, body.Email, string(hashed))

	if err != nil {
		http.Error(w, "User already exists or DB error", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&body)

	user, err := models.GetUserByEmail(body.Email)
	if err != nil {
		http.Error(w, "Well you have inputted invalid credentials my friend", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		http.Error(w, "Wrong Password Brochahco", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenStr, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	json.NewEncoder(w).Encode(map[string]string{"token": tokenStr})
}
