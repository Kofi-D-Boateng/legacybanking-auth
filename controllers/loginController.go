package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kofi-D-Boateng/legacybanking-auth/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginRequest struct {
		Email    string
		Password string
	}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&loginRequest)

	if err != nil {
		fmt.Println("[ERROR]: There was an error retrieving email and password variables....: ", err)
		w.WriteHeader(400)
		return
	}

	query := "SELECT * FROM customer WHERE email = $1;"
	stmt, err := utils.DatabaseConn.PrepareContext(r.Context(), query)

	if err != nil {
		fmt.Printf("[ERROR]: Error preparing PrepareContext statement....:%v", err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(loginRequest.Email)

	if err != nil {
		fmt.Printf("[ERROR]: Error querying database....: %v", err)
	}

	defer rows.Close()

	var returnedValue struct {
		email    string
		password string
	}

	for rows.Next() {
		if err := rows.Scan(&returnedValue); err != nil {
			fmt.Printf("[ERROR]: Error mapping struct to retieve email and password...: %v", err)
		}
	}

	mismatchPassword := bcrypt.CompareHashAndPassword([]byte(loginRequest.Password), []byte(returnedValue.password))

	if mismatchPassword != nil {
		fmt.Println("[ERROR]: Password does not match hashed password")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	jwtToken := utils.CreateJwt(returnedValue.email)

	var returningValues struct {
		Email     string
		AuthToken string
		ApiKey    string
	}

	returningValues.Email = returnedValue.email
	returningValues.AuthToken = jwtToken
	returningValues.ApiKey = uuid.New().String()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(returningValues)

}

func LoginEmployee(w http.ResponseWriter, r *http.Request) {

}
