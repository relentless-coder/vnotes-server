package notebooks

import "github.com/gorilla/mux"

func Router(r *mux.Router) {
	s := r.PathPrefix("/notebooks").Subrouter()
	s.HandleFunc("/", getNotebooks).Methods("GET")
	s.HandleFunc("/", createNotebook).Methods("POST")
	s.HandleFunc("/{id}", updateNotebook).Methods("PUT")
}
