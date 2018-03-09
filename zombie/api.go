package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
	r.HandleFunc("/drivers/{id}", s.getDriver).Methods("GET")
	http.Handle("/", r)

	return r
}

func (s *API) getHome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Zombie"))
	handleError(err)
}

func (s *API) getDriver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	driverID, err := strconv.Atoi(mux.Vars(r)["id"])
	handleError(err)

	driver := Driver{ID: driverID, Zombie: true}
	driverJSON, err := json.Marshal(driver)
	handleError(err)

	_, err = w.Write(driverJSON)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
