package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Kofi-D-Boateng/legacybanking-auth/utils"
)

func AuthenticateUser(paylod json.RawMessage) (utils.Response,error) {

	var authToken string
	err:= json.Unmarshal(paylod,&authToken)
	if err != nil{
		return utils.Response{StatusCode: http.StatusUnauthorized,Body: []byte("")},err
	}
	email, jwtErr := utils.VerifyJwt(authToken)
	if jwtErr != nil {
		return utils.Response{StatusCode: http.StatusUnauthorized,Body: []byte("")},nil
	}
	t,_ := json.Marshal(email)
	return utils.Response{StatusCode: http.StatusOK,Body:t },nil
}

func AuthenticateEmployee(paylod json.RawMessage)(utils.Response,error) {
	var authToken string
	err:= json.Unmarshal(paylod,&authToken)
	if err != nil{
		return utils.Response{StatusCode: http.StatusUnauthorized,Body: []byte("")},err
	}
	email, jwtErr := utils.VerifyJwt(authToken)
	if jwtErr != nil {
		return utils.Response{StatusCode: http.StatusUnauthorized,Body: []byte("")},nil
	}
	return utils.Response{StatusCode: http.StatusOK,Body: []byte(email)},nil
}
