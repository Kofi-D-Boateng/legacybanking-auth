package utils

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DatabaseConn *sqlx.DB

func ConnectSQLDatabase(driverName string, connStr string) {

	conn, err := sqlx.Open(driverName, connStr)

	if err != nil {
		log.Fatalf("[ERROR]: Connection to %s database using connection string: %s did not work...", driverName, connStr)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatalf("[ERROR]: Cannot ping db: %v", err)
		panic(err)
	}

	DatabaseConn = conn
}
