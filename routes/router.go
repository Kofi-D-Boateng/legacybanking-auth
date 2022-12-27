package routes

import (
	"fmt"
	"os"

	auth "github.com/Kofi-D-Boateng/legacy-banking-api/controllers/authentication"
	bank "github.com/Kofi-D-Boateng/legacy-banking-api/controllers/bank"
	billing "github.com/Kofi-D-Boateng/legacy-banking-api/controllers/billing"
	notification "github.com/Kofi-D-Boateng/legacy-banking-api/controllers/notifications"
	profile "github.com/Kofi-D-Boateng/legacy-banking-api/controllers/profile"
	registration "github.com/Kofi-D-Boateng/legacy-banking-api/controllers/registration"
	security "github.com/Kofi-D-Boateng/legacy-banking-api/controllers/security"
	transaction "github.com/Kofi-D-Boateng/legacy-banking-api/controllers/transaction"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	apiVersion := os.Getenv("API_VERSION")

	// GET REQUESTS
	profileInfoUri := fmt.Sprintf("/%s/authentication/profile",apiVersion);
	confirmInfoUri := fmt.Sprintf("/%s/authentication/confirm-account",apiVersion)
	refreshTokenUri := fmt.Sprintf("/%s/authentication/get-refresh-token",apiVersion)
	bankInfoUri := fmt.Sprintf("/%s/bank/info",apiVersion)
	router.HandleFunc(profileInfoUri,profile.GetProfileInfo).Methods("GET")
	router.HandleFunc(confirmInfoUri,auth.ConfirmCustomerAccount).Methods("GET")
	router.HandleFunc(refreshTokenUri,auth.CreateRefreshToken).Methods("GET")
	router.HandleFunc(bankInfoUri,bank.GetBankInfo).Methods("GET")

	// POST REQUESTS
	registrationUri := fmt.Sprintf("/%s/authentication/admin/registration",apiVersion)
	loginUri := fmt.Sprintf("/%s/authentication/login",apiVersion)
	verificationLinkUri := fmt.Sprintf("/%s/authentication/new-verifcation-link",apiVersion)
	router.HandleFunc(registrationUri,registration.RegisterUser).Methods("POST")
	router.HandleFunc(loginUri,auth.LoginUser).Methods("POST")
	router.HandleFunc(verificationLinkUri,auth.SendNewVerificationLink).Methods("POST")

	// PUT REQUESTS
	securityUpdateUri := fmt.Sprintf("/%s/authentication/profile/update-account-security",apiVersion)
	billingUpdateUri := fmt.Sprintf("/%s/authentication/update-billing-type",apiVersion)
	transactionUri := fmt.Sprintf("/%s/authentication/process-transaction",apiVersion)
	notificationUri := fmt.Sprintf("/%s/authentication/send-notifcation",apiVersion)
	router.HandleFunc(securityUpdateUri,security.UpdateAccountSecurity).Methods("PUT")
	router.HandleFunc(billingUpdateUri,billing.UpdateBillingType).Methods("PUT")
	router.HandleFunc(transactionUri,transaction.ProcessTransaction).Methods("PUT")
	router.HandleFunc(notificationUri,notification.UpdateNotification).Methods("PUT")

	

	return router
}