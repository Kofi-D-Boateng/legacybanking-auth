package controllers

import (
	"fmt"
	"net/http"

	"github.com/Kofi-D-Boateng/legacybanking-auth/utils"
)

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("authorization")

	if authToken == "" {
		fmt.Print("Authorization token is not present")

		w.WriteHeader(http.StatusUnauthorized)
	}

	_, err := utils.VerifyJwt(authToken)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}

	w.WriteHeader(http.StatusOK)
}

func AuthenticateEmployee(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("authorization")

	if authToken == "" {
		fmt.Print("Authorization token is not present")

		w.WriteHeader(http.StatusUnauthorized)
	}
	// FOR NOW WE WILL RETURN TRUE
	// THIS CAN BE CURTAILED TO USE
	// JWT OR LDAP

	w.WriteHeader(http.StatusOK)
}
