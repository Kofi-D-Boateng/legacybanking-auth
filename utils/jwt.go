package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateJwt(email string) string {
	secret := os.Getenv("JWT_SECRET")
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":   "legacybanking-auth.com",
		"exp":   time.Now().Add(time.Minute * 10).Unix(),
		"email": email,
	})

	signedToken, err := jwtToken.SignedString([]byte(secret))

	if err != nil {
		log.Fatalf("[ERROR]: Error signing tokens....: %v", err)
	}
	fmt.Printf("JWT: %v", signedToken)
	return signedToken
}

func VerifyJwt(hashedToken string) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.Parse(hashedToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return "Failed to parse token", err
	}

	if !token.Valid {
		return "Invalid token used", fmt.Errorf("result:%v ", token.Valid)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("failed to cast claims to map")
		return "failed to cast claims to map", fmt.Errorf("%v", ok)
	}

	email, ok := claims["email"].(string)
	if !ok {
		return "failed to cast a claim key to the concrete type string", fmt.Errorf("%v", ok)
	}

	return email, nil
}
