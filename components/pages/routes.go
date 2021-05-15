package pages

import "github.com/gorilla/mux"

func Routes(r *mux.Router) {
	s := r.PathPrefix("/notebooks/{notebooks_id}/pages").Subrouter()
	s.HandleFunc("/", getPages).Methods("GET")
	s.HandleFunc("/", createPage).Methods("POST")
	s.HandleFunc("/{id}", updatePage).Methods("PUT")
}
