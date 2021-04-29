package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func init() {
	userName := os.Getenv("APP_DB_USERNAME")
	dbName := os.Getenv("APP_DB_NAME")
	connectionString := fmt.Sprintf("user=%s dbname=%s sslmode=disable", userName, dbName)
	fmt.Printf(connectionString)
	var err error
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Panic(err)
	}
}
