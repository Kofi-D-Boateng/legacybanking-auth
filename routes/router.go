package routes

import (
	"fmt"
	"os"

	"github.com/Kofi-D-Boateng/legacybanking-auth/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	apiVersion := os.Getenv("API_VERSION")

	authCustomerUri := fmt.Sprintf("/%s/auth/authenticate-customer", apiVersion)
	authEmployeeUri := fmt.Sprintf("/%s/auth/admin/authenticate-employee", apiVersion)
	refreshTokenUri := fmt.Sprintf("/%s/auth/get-refresh-token", apiVersion)
	confirmUserUri := fmt.Sprintf("/%s/auth/confirm-user-account", apiVersion)
	logoutUri := fmt.Sprintf("/%s/auth/logout", apiVersion)
	router.HandleFunc(authCustomerUri, controllers.AuthenticateUser).Methods("GET")
	router.HandleFunc(authEmployeeUri, controllers.AuthenticateEmployee).Methods("GET")
	router.HandleFunc(logoutUri, controllers.LogoutUser).Methods("GET")
	router.HandleFunc(refreshTokenUri, controllers.CreateRefreshToken).Methods("GET")
	router.HandleFunc(confirmUserUri, controllers.ConfirmUser).Methods("GET")

	// POST REQUESTS
	registrationUri := fmt.Sprintf("/%s/auth/admin/registration", apiVersion)
	loginCustomerUri := fmt.Sprintf("/%s/auth/login-customer", apiVersion)
	loginEmplyeeUri := fmt.Sprintf("/%s/auth/login-employee", apiVersion)
	router.HandleFunc(registrationUri, controllers.RegisterUser).Methods("POST")
	router.HandleFunc(loginCustomerUri, controllers.LoginUser).Methods("POST")
	router.HandleFunc(loginEmplyeeUri, controllers.LoginEmployee).Methods("POST")

	return router
}
