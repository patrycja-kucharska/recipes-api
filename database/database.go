package database

import (
	"fmt"
	"os"
	"strconv"

	// pq is used as a driver (behind the courtains) by sqlx (defined in Connect statement with "postgres" keyword)
	// and as a driver for loggerAdapter
	pq "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
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

	loggerAdapter := zerologadapter.New(zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}))
	db.DB = sqldblogger.OpenDriver(dbConfig, &pq.Driver{}, loggerAdapter)

	// defer db.Close()
}
