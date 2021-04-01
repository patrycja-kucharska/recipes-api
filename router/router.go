package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Router() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/ingredient", postIngredient).Methods("POST")
	r.HandleFunc("/ingredient", findIngredients).Methods("GET").Queries("name", "{name}")
	r.HandleFunc("/ingredient/{id}", getIngredient).Methods("GET")
	r.HandleFunc("/ingredients", getAllIngredients).Methods("GET")

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	log.Fatal(http.ListenAndServe(":8080", loggedRouter))
}
