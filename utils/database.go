package utils

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DatabaseConn *sql.Conn

func ConnectSQLDatabase(driverName string, connStr string) {
	DatabaseConn, err := sql.Open(driverName, connStr)

	if err != nil {
		log.Fatalf("[ERROR]: Connection to %s database using connection string: %s did not work...", driverName, connStr)
	}

	defer DatabaseConn.Close()
}
