package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patrycja-kucharska/recipes-api/database"
	"github.com/patrycja-kucharska/recipes-api/structs"
)

func postIngredient(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var newIngredient structs.Ingredient
	json.Unmarshal(body, &newIngredient)
	id, err := database.AddIngredient(newIngredient)

	if err != nil {
		createResponse(w, "error", err.Error(), nil, http.StatusUnprocessableEntity)
		return
	}
	reason := "Successfully created Ingredient/" + id
	createResponse(w, "success", reason, nil, http.StatusCreated)
}

func getIngredient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ing, err := database.SelectIngredient(structs.Ingredient{Id: vars["id"]})
	if err != nil {
		createResponse(w, "error", err.Error(), nil, http.StatusInternalServerError)
		return
	}
	createResponse(w, "success", "", ing, http.StatusOK)
}

func getAllIngredients(w http.ResponseWriter, r *http.Request) {
	ing, err := database.ListIngredients()
	if err != nil {
		createResponse(w, "error", err.Error(), nil, http.StatusInternalServerError)
		return
	}
	createResponse(w, "success", "", ing, http.StatusOK)
}

func findIngredients(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ing, err := database.SelectIngredientsWhere(vars)
	if err != nil {
		createResponse(w, "error", err.Error(), nil, http.StatusInternalServerError)
		return
	}

	createResponse(w, "success", "", ing, http.StatusOK)
}

func createResponse(w http.ResponseWriter, code, reason string, body interface{}, statusCode int) {
	resp := structs.Response{}
	resp.Code = code
	if reason != "" {
		resp.Reason = reason
	}
	if body != nil {
		resp.Body = body
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}
