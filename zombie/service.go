package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Service contains the HTTP router and serves the API
type Service struct {
	Router *mux.Router
}

// Initialise creates and sets the API router
func (s *Service) Initialise() {
	s.Router = mux.NewRouter()
	s.Router.HandleFunc("/", s.GetHome).Methods("GET")
	http.Handle("/", s.Router)
}

// Start will start the api on a given port
func (s *Service) Start(port string) {
	log.Fatal(http.ListenAndServe(port, s.Router))
}

// GetHome displays a welcome message
func (s *Service) GetHome(w http.ResponseWriter, r *http.Request) {
	//	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Welcome to Zombie")
}
