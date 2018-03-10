package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// API contains the HTTP router and serves the API.
type API struct {
	router *mux.Router
	store  *Store
}

// Serve will start the api on a given port.
func (a *API) Serve(port string) {
	log.Fatal(http.ListenAndServe(port, a.NewRouter()))
}

// ConnectStore attaches the API to the redis backed store.
func (a *API) ConnectStore(s *Store) {
	a.store = s
}

// NewRouter creates and returns a pointer to an API router.
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

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	handleAPIError(err)

	log.Printf("DriverID: %v\n", DriverID(id))

	locations := a.store.GetLocations(DriverID(id))

	locationsJSON, err := json.Marshal(locations)
	handleAPIError(err)

	_, err = w.Write(locationsJSON)
	handleAPIError(err)
}

func handleAPIError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
