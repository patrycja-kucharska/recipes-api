package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	// pq is used as a driver (behind the courtains) by sqlx (defined in Connect statement with "postgres" keyword)
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env", ".base.env")
	connectToDB()
	router()
}

func router() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/ingredient", postIngredient).Methods("POST")
	r.HandleFunc("/ingredient", findIngredients).Methods("GET").Queries("name", "{name}")
	r.HandleFunc("/ingredient/{id}", getIngredient).Methods("GET")
	r.HandleFunc("/ingredients", getAllIngredients).Methods("GET")

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	log.Fatal(http.ListenAndServe(":8080", loggedRouter))
}
