package database

import (
	"fmt"
	"log"
	"strings"

	"github.com/patrycja-kucharska/recipes-api/structs"
)

func AddIngredient(ingredient structs.Ingredient) (string, error) {
	namedStmt, err := db.PrepareNamed(`INSERT INTO ingredients(name, unit_name) VALUES(:name, :unit_name) RETURNING id`)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
		return "", err
	}
	namedStmt.QueryRowx(ingredient).StructScan(&ingredient)
	return ingredient.Id, nil
}

func SelectIngredient(interf structs.Ingredient) (structs.Ingredient, error) {
	var ingredient structs.Ingredient

	namedStmt, err := db.PrepareNamed(`SELECT *, 'Ingredient' as type FROM ingredients WHERE id=:id`)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
		return ingredient, err
	}
	err2 := namedStmt.Get(&ingredient, interf)

	return ingredient, err2
}

func SelectIngredientsWhere(params map[string]string) ([]structs.Ingredient, error) {
	var ingredients []structs.Ingredient

	var qString []string
	for key, value := range params {
		qString = append(qString, fmt.Sprintf(`%s='%s'`, key, value))
	}

	rows, err := db.Queryx(`SELECT *, 'Ingredient' as type FROM ingredients WHERE ` + strings.Join(qString, " AND "))

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
		return ingredients, err
	}

	for rows.Next() {
		var oneIngredient structs.Ingredient
		rows.StructScan(&oneIngredient)
		ingredients = append(ingredients, oneIngredient)
	}

	return ingredients, nil
}

func ListIngredients() ([]structs.Ingredient, error) {
	var ingredients []structs.Ingredient

	rows, err := db.Queryx(`SELECT * , 'Ingredient' as type FROM ingredients`)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
		return ingredients, err
	}

	for rows.Next() {
		var oneIngredient structs.Ingredient
		rows.StructScan(&oneIngredient)
		ingredients = append(ingredients, oneIngredient)
	}

	return ingredients, nil
}
