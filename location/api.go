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
func (a *API) Serve(port string) {
	log.Fatal(http.ListenAndServe(port, a.NewRouter()))
}

// NewRouter creates and returns a pointer to an API router
func (a *API) NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", a.getHome).Methods("GET")
	r.HandleFunc("/drivers/{id}/coordinates", a.getLocations).Methods("GET")
	http.Handle("/", r)

	return r
}

func (a *API) getHome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Location"))
	handleAPIError(err)
}

func (a *API) getLocations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	locationsJSON, err := json.Marshal([]Location{
		Location{
			Latitude:  42,
			Longitude: 2.3,
			UpdatedAt: "YYYY-MM-DDTHH:MM:SSZ",
		},
		Location{
			Latitude:  42.1,
			Longitude: 2.32,
			UpdatedAt: "YYYY-MM-DDTHH:MM:SSZ",
		},
	})
	handleAPIError(err)

	_, err = w.Write(locationsJSON)
	handleAPIError(err)
}

func handleAPIError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
