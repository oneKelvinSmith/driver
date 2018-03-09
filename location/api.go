package main

import (
	"encoding/json"
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
	r.HandleFunc("/drivers/{id}/coordinates", s.getLocations).Methods("GET")
	http.Handle("/", r)

	return r
}

func (s *API) getHome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Location"))
	handleError(err)
}

func (s *API) getLocations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	locationsJSON, err := json.Marshal([]Location{
		Location{
			Latitude:  42,
			Longitude: 2.3,
			UpdatedAt: "YYYY-MM-DDTHH:MM:SSZ",
		},
		Location{Latitude: 42.1,
			Longitude: 2.32,
			UpdatedAt: "YYYY-MM-DDTHH:MM:SSZ",
		},
	})
	handleError(err)

	_, err = w.Write(locationsJSON)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
