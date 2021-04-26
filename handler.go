package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) getPages(w http.ResponseWriter, r *http.Request) {
	pages, err := getPages(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusAccepted, pages)
}

func (a *App) createPage(w http.ResponseWriter, r *http.Request) {
	var p pages
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	p.CreatedAt = time.Now().String()
	p.UpdatedAt = time.Now().String()
	fmt.Print(p)
	if err := p.createPage(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, p)
}

func (a *App) updatePage(w http.ResponseWriter, r *http.Request) {
	var p pages
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid page id")
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	p.ID = id
	p.UpdatedAt = time.Now().String()
	if err := p.updatePage(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, p)
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/pages", a.createPage).Methods("POST")
	a.Router.HandleFunc("/pages", a.getPages).Methods("GET")
	a.Router.HandleFunc("/pages/{id:[0-9]+}", a.updatePage).Methods("PUT")
}
