package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	// pq is used as a driver (behind the courtains) by sqlx (defined in Connect statement with "postgres" keyword)
	_ "github.com/lib/pq"
)

func main() {
	handleRequests()
}

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	connectToDB()
	r.HandleFunc("/ingredient", postIngredient).Methods("POST")
	r.HandleFunc("/ingredient/{id}", getIngredient).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
