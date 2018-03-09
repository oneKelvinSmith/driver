package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// API contains the HTTP router and serves the API
type API struct {
	router *mux.Router
}

// Serve will start the api on a given port
func (s *API) Serve(port string) {
	log.Fatal(http.ListenAndServe(port, s.NewRouter()))
}

// NewRouter creates and returns a pointer to an API router
func (s *API) NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", s.getHome).Methods("GET")
	http.Handle("/", r)

	return r
}

func (s *API) getHome(w http.ResponseWriter, r *http.Request) {
	//	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Welcome to Zombie")
}
