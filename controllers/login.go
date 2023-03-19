package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kofi-D-Boateng/legacybanking-auth/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(payload json.RawMessage)(utils.Response,error) {

	var loginRequest struct {
		Email    string `json:"email"`
		Password string	`json:"password"`
	}
	
	var returnedValue struct {
		Email       string
		Password    string
		IsActivated bool
	}

	err := json.Unmarshal(payload,&loginRequest)

	if err != nil{
		return utils.Response{StatusCode: http.StatusUnauthorized,Body:[]byte("")},err
	}

	fmt.Printf("LoginRequest --> %v",loginRequest)

	query := "SELECT email,password,is_activated FROM customer WHERE email = $1;"
	queryErr := utils.DatabaseConn.QueryRow(query, loginRequest.Email).Scan(&returnedValue.Email, &returnedValue.Password, &returnedValue.IsActivated)

	if queryErr != nil {
		fmt.Printf("[ERROR]: Error querying database....: %v", queryErr)
		return utils.Response{StatusCode: http.StatusUnauthorized,Body: []byte("")},queryErr
	}
	
	fmt.Printf("Queried Data --> %v",returnedValue)

	mismatchPassword := bcrypt.CompareHashAndPassword([]byte(returnedValue.Password), []byte(loginRequest.Password))

	if mismatchPassword != nil {
		fmt.Printf("[ERROR]: Password does not match hashed password: %v\n", mismatchPassword)
		return utils.Response{StatusCode: http.StatusUnauthorized,Body:[]byte("")},mismatchPassword
	}


	jwtToken, expiresAt := utils.CreateJwt(returnedValue.Email)

	fmt.Println("Successful JWT creation....")

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

	rv,_ := json.Marshal(returningValues)
	return utils.Response{StatusCode: http.StatusOK,Body: rv},nil

}

func LoginEmployee(payload interface{})(utils.Response,error) {
	return utils.Response{},nil
}
