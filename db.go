package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func connectToDB() {
	dbHost := os.Getenv("DB_HOST")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	var dbConfig string = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbName)
	db = sqlx.MustConnect("postgres", dbConfig)
	// defer db.Close()
}

func addIngredient(ingredient Ingredient) (string, error) {
	namedStmt, err := db.PrepareNamed(`INSERT INTO ingredients(name, unit_name) VALUES(:name, :unit_name) RETURNING id`)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
		return "", err
	}
	namedStmt.QueryRowx(ingredient).StructScan(&ingredient)
	return ingredient.Id, nil
}

func selectIngredient(interf Ingredient) (Ingredient, error) {
	var ingredient Ingredient

	namedStmt, err := db.PrepareNamed(`SELECT *, 'Ingredient' as type FROM ingredients WHERE id=:id`)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
		return ingredient, err
	}
	err2 := namedStmt.Get(&ingredient, interf)

	return ingredient, err2
}

func listIngredients() ([]Ingredient, error) {
	var ingredients []Ingredient

	rows, err := db.Queryx(`SELECT * , 'Ingredient' as type FROM ingredients`)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
		return ingredients, err
	}

	for rows.Next() {
		var oneIngredient Ingredient
		rows.StructScan(&oneIngredient)
		ingredients = append(ingredients, oneIngredient)
	}

	return ingredients, nil
}
