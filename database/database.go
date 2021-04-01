package database

import (
	"fmt"
	"os"
	"strconv"

	// pq is used as a driver (behind the courtains) by sqlx (defined in Connect statement with "postgres" keyword)
	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func ConnectToDB() {
	dbHost := os.Getenv("DB_HOST")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	var dbConfig string = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbName)
	db = sqlx.MustConnect("postgres", dbConfig)
	// defer db.Close()
}
