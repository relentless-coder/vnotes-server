package notebooks

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	c "github.com/relentless-coder/vnotes-server/config"
	h "github.com/relentless-coder/vnotes-server/helpers"
)

func getNotebooks(w http.ResponseWriter, r *http.Request) {
	var n notebooks
	nbs, err := n.get(c.DB)
	if err != nil {
		h.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	h.RespondWithJSON(w, http.StatusOK, nbs)
}

func createNotebook(w http.ResponseWriter, r *http.Request) {
	var n notebooks
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&n); err != nil {
		h.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	n.CreatedAt = time.Now().String()
	n.UpdatedAt = time.Now().String()
	if err := n.create(c.DB); err != nil {
		h.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	h.RespondWithJSON(w, http.StatusCreated, n)
}

func updateNotebook(w http.ResponseWriter, r *http.Request) {
	var n notebooks
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.RespondWithError(w, http.StatusBadRequest, "Invalid page id")
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&n); err != nil {
		h.RespondWithError(w, http.StatusInsufficientStorage, err.Error())
	}
	n.Id = id
	n.UpdatedAt = time.Now().String()
	if err := n.update(c.DB); err != nil {
		h.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	h.RespondWithJSON(w, http.StatusOK, n)
}
