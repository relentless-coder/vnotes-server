package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/relentless-coder/vnotes-server/components"
)

func main() {
	var router *mux.Router
	router = components.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8080", router))
}
