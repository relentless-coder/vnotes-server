package pages

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	c "github.com/relentless-coder/vnotes-server/config"
	h "github.com/relentless-coder/vnotes-server/helpers"
)

func getPages(w http.ResponseWriter, r *http.Request) {
	var p pages
	pages, err := p.get(c.DB)
	if err != nil {
		h.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.RespondWithJSON(w, http.StatusAccepted, pages)
}

func createPage(w http.ResponseWriter, r *http.Request) {
	var p pages
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		h.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	p.CreatedAt = time.Now().String()
	p.UpdatedAt = time.Now().String()
	fmt.Print(p)
	if err := p.create(c.DB); err != nil {
		h.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.RespondWithJSON(w, http.StatusCreated, p)
}

func updatePage(w http.ResponseWriter, r *http.Request) {
	var p pages
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.RespondWithError(w, http.StatusBadRequest, "Invalid page id")
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		h.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	p.Id = id
	p.UpdatedAt = time.Now().String()
	if err := p.update(c.DB); err != nil {
		h.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.RespondWithJSON(w, http.StatusOK, p)
}
