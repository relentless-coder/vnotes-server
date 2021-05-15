package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func init() {
	SetConfig()
	dbConfig, err := GetDBVars()
	if err != nil {
		log.Fatal("err fetching data config %v\n", err)
	}
	connectionString := fmt.Sprintf("user=%s dbname=%s password=%s  sslmode=disable", dbConfig.User, dbConfig.Dbname, dbConfig.Password)
	fmt.Printf(connectionString)
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Panic(err)
	}
}
