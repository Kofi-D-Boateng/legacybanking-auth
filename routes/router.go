package routes

import (
	"fmt"
	"os"

	"github.com/Kofi-D-Boateng/legacybanking-auth/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	apiVersion := os.Getenv("API_VERSION")

	authenticationUri := fmt.Sprintf("/%s/auth/authenticate-user", apiVersion)
	refreshTokenUri := fmt.Sprintf("/%s/auth/get-refresh-token", apiVersion)
	confirmUserUri := fmt.Sprintf("/%s/auth/confirm-user-account", apiVersion)
	logoutUri := fmt.Sprintf("/%s/auth/logout", apiVersion)
	router.HandleFunc(authenticationUri, controllers.AuthenticateUser).Methods("POST")
	router.HandleFunc(logoutUri, controllers.LogoutUser).Methods("GET")
	router.HandleFunc(refreshTokenUri, controllers.CreateRefreshToken).Methods("GET")
	router.HandleFunc(confirmUserUri, controllers.ConfirmUser).Methods("GET")

	// POST REQUESTS
	registrationUri := fmt.Sprintf("/%s/auth/admin/registration", apiVersion)
	loginUri := fmt.Sprintf("/%s/auth/login", apiVersion)
	router.HandleFunc(registrationUri, controllers.RegisterUser).Methods("POST")
	router.HandleFunc(loginUri, controllers.LoginUser).Methods("POST")

	return router
}
