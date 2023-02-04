package controllers

import (
	"fmt"
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("authorization")

	if authToken == "" {
		fmt.Print("Authorization token is not present")

		w.WriteHeader(401)
	}

	secret := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.Parse(authToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		fmt.Println("Failed to parse token: ", err)
		w.WriteHeader(401)
	}

	if !token.Valid {
		fmt.Println("Invalid token used")
		w.WriteHeader(401)
	}
}
