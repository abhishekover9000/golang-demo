package utils

import (
	"encoding/json"
	"net/http"
	"os"

	"../models"
	jwt "github.com/dgrijalva/jwt-go"
)

// RespondWithError yo this does something
func RespondWithError(w http.ResponseWriter, status int, error models.Error) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(error)
}

// RespondwithJSON yo this does something
func RespondWithJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// GenerateToken responds with things
func GenerateToken(user models.User) (string, error) {
	var err error
	secret := os.Getenv("SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})

	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, err
}
