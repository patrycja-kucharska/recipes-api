package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func postIngredient(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var newIngredient Ingredient
	json.Unmarshal(body, &newIngredient)
	id, err := addIngredient(newIngredient)

	if err != nil {
		createResponse(w, "error", err.Error(), nil, http.StatusUnprocessableEntity)
		return
	}
	reason := "Successfully created Ingredient/" + id
	createResponse(w, "success", reason, nil, http.StatusCreated)
}

func getIngredient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ing, err := selectIngredient(Ingredient{Id: vars["id"]})
	if err != nil {
		createResponse(w, "error", err.Error(), nil, http.StatusInternalServerError)
		return
	}
	createResponse(w, "success", "", ing, http.StatusOK)
}

func getAllIngredients(w http.ResponseWriter, r *http.Request) {
	ing, err := listIngredients()
	if err != nil {
		createResponse(w, "error", err.Error(), nil, http.StatusInternalServerError)
		return
	}
	createResponse(w, "success", "", ing, http.StatusOK)
}

func findIngredients(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ing, err := selectIngredientsWhere(vars)
	if err != nil {
		createResponse(w, "error", err.Error(), nil, http.StatusInternalServerError)
		return
	}

	createResponse(w, "success", "", ing, http.StatusOK)
}

func createResponse(w http.ResponseWriter, code, reason string, body interface{}, statusCode int) {
	resp := Response{}
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
