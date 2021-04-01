package main

import (
	"github.com/joho/godotenv"

	// pq is used as a driver (behind the courtains) by sqlx (defined in Connect statement with "postgres" keyword)
	_ "github.com/lib/pq"

	"github.com/patrycja-kucharska/recipes-api/database"
	"github.com/patrycja-kucharska/recipes-api/router"
)

func main() {
	godotenv.Load(".env", ".base.env")
	database.ConnectToDB()
	router.Router()
}
