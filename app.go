package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbName string) {
	connectionString := fmt.Sprintf("user=%s dbname=%s sslmode=disable", user, dbName)
	fmt.Printf(connectionString)
	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	res, err := a.DB.Query("select table_catalog,table_name,table_schema from information_schema.tables where table_schema NOT IN ('pg_catalog', 'information_schema')")
	if err != nil {
		log.Panic(err)
	}
	type Row struct {
		Catalog string `json:"table_catalog"`
		Name    string `json:"table_name"`
		Schema  string `json:"table_schema"`
	}
	for res.Next() {
		var table Row
		if err := res.Scan(&table.Catalog, &table.Name, &table.Schema); err != nil {
			log.Fatal(err)
		}
		fmt.Println(table)
	}
	if err != nil {
		log.Panic(err)
	}
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(add string) {
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}
