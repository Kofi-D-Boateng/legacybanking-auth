package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Kofi-D-Boateng/legacybanking-auth/utils"
)

func ConfirmUser(payload json.RawMessage)(utils.Response,error) {
	var token string
	var expiresAt string

	err := json.Unmarshal(payload,&token)

	if err != nil {
		return utils.Response{StatusCode: http.StatusUnauthorized,Body:[]byte("")},err
	}
	query := "SELECT expires_at FROM verification_token WHERE token = $1;"
	queryErr := utils.DatabaseConn.QueryRow(query, token).Scan(&expiresAt)

	if queryErr != nil {
		fmt.Printf("[ERROR]: Error querying database....: %v", queryErr)
		return utils.Response{StatusCode: http.StatusUnauthorized,Body: []byte("")},queryErr
	}
	
	fmt.Printf("Queried Timestamp Data --> %v",expiresAt)

	timestamp, err := time.Parse("2006-01-02 15:04:05",expiresAt)
	timeElapsed := timestamp.After(time.Now())

	if err != nil {
		return utils.Response{StatusCode: http.StatusUnauthorized,Body:[]byte("")},err
	}

	if timeElapsed {
		return utils.Response{StatusCode: http.StatusUnauthorized,Body: []byte("")}, errors.New("confirmation token is expired")
	}

	return utils.Response{StatusCode: http.StatusOK,Body: []byte("")},nil
}
