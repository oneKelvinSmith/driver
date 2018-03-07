package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetDriver(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Driver: %v\n", vars["id"])
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/drivers/{id}", GetDriver).Methods("GET")
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":3000", router))
}
