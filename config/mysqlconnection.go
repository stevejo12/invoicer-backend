package config

import (
	"database/sql"
	"fmt"
	"os"
)

var DB *sql.DB

// dsn => data source name
// getting the database dsn
func dsn() string {
	return (fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOSTNAME"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	))
}

func InitializeConnectionToMySQLDatabase() error {
	var err error

	DB, err = sql.Open("mysql", dsn())

	if err != nil {
		return err
	}

	return DB.Ping()
}
