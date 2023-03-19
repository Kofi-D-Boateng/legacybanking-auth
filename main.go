package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Kofi-D-Boateng/legacybanking-auth/controllers"
	"github.com/Kofi-D-Boateng/legacybanking-auth/utils"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("ENV") == "dev" {
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

	connStr := os.Getenv("DB_CONN")
	driverName := os.Getenv("DB_DRIVER")

	utils.ConnectSQLDatabase(driverName, connStr)
	defer utils.DatabaseConn.Close()
	lambda.Start(handler)
}


func handler(ctx context.Context, req utils.Request) (utils.Response,error){

		fmt.Printf("Request --> %v\n",req)
		
		switch req.Function{
			case "authenticateUser":
				return controllers.AuthenticateUser(req.Payload)
			case "authenticateEmployee":
				return controllers.AuthenticateEmployee(req.Payload)
			case "loginUser":
				return controllers.LoginUser(req.Payload)
			case "loginEmployee":
				return controllers.LoginEmployee(req.Payload)	
			case "confirmUser":
				return controllers.ConfirmUser(req.Payload)
			case "getRefreshToken":
				return controllers.CreateRefreshToken(req.Payload)
			default:
				return utils.Response{},errors.New("unknown function")				
		}
	
	}