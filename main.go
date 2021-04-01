package main

import (
	"github.com/joho/godotenv"

	"github.com/patrycja-kucharska/recipes-api/database"
	"github.com/patrycja-kucharska/recipes-api/router"
)

func main() {
	godotenv.Load(".env", ".base.env")
	database.ConnectToDB()
	router.Router()
}
