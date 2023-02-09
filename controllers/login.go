package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kofi-D-Boateng/legacybanking-auth/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var loginRequest struct {
		Email    string
		Password string
	}

	var returnedValue struct {
		Email       string
		Password    string
		IsActivated bool
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginRequest)

	if err != nil {
		fmt.Println("[ERROR]: There was an error retrieving email and password variables....: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	query := "SELECT email,password,is_activated FROM customer WHERE email = $1;"
	queryErr := utils.DatabaseConn.QueryRow(query, loginRequest.Email).Scan(&returnedValue.Email, &returnedValue.Password, &returnedValue.IsActivated)

	if queryErr != nil {
		fmt.Printf("[ERROR]: Error querying database....: %v", queryErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	mismatchPassword := bcrypt.CompareHashAndPassword([]byte(returnedValue.Password), []byte(loginRequest.Password))

	if mismatchPassword != nil {
		fmt.Printf("[ERROR]: Password does not match hashed password: %v\n", mismatchPassword)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	jwtToken, expiresAt := utils.CreateJwt(returnedValue.Email)

	var returningValues struct {
		AuthToken       string
		ApiKey          string
		TokenExpiration int64
		IsActivated     bool
	}

	returningValues.AuthToken = jwtToken
	returningValues.ApiKey = uuid.New().String()
	returningValues.TokenExpiration = expiresAt
	returningValues.IsActivated = returnedValue.IsActivated

	redisErr := utils.RedisClient.Set(context.Background(), returningValues.ApiKey, returnedValue.Email, 0).Err()
	if redisErr != nil {
		panic(redisErr)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(returningValues)

}

func LoginEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
