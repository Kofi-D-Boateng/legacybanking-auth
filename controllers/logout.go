package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Kofi-D-Boateng/legacybanking-auth/utils"
)

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authToken := r.Header.Get("authorization")
	apiKey := r.URL.Query().Get("apiKey")

	if authToken == "" || apiKey == "" {
		fmt.Print("Authorization token or apiKey is not present")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, verErr := utils.VerifyJwt(authToken)

	if verErr != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	redisErr := utils.RedisClient.Del(context.TODO(), apiKey).Err()
	if redisErr != nil {
		log.Fatalln(redisErr)
		w.WriteHeader(http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusOK)
}
