package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Kofi-D-Boateng/legacybanking-auth/routes"
	"github.com/Kofi-D-Boateng/legacybanking-auth/utils"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("ENV") == "dev" || os.Getenv("ENV") == "" {
		_, file, _, ok := runtime.Caller(0)
		basePath := filepath.Dir(file)
		fmt.Println(file)
		fmt.Println(basePath)

		if !ok {
			log.Fatalf("Unable to find file path: %v", file)
		}

		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
	}
}

func main() {
	router := routes.Router()
	port := os.Getenv("PORT")
	redisAddr := os.Getenv("REDIS_ADDR")
	connStr := os.Getenv("DB_CONN")
	driverName := os.Getenv("DB_DRIVER")
	messageBrokerAddr := os.Getenv("MESSAGE_BROKER_ADDR")

	allowedHeaders := handlers.AllowedHeaders([]string{"Origin", "Content-Type", "Accept", "authorization", "x-forwarded-for", "User-Agent"})
	allowedOrigins := handlers.AllowedOrigins([]string{os.Getenv("ORIGINS")})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT"})

	utils.ConnectRedis(redisAddr, "", 0)
	utils.ConnectMessageBroker(messageBrokerAddr)
	utils.ConnectSQLDatabase(driverName, connStr)
	fmt.Printf("Server listening at port:%v \n", port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)(router)))

}
