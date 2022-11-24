package controllers

import (
	"encoding/json"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	"net/http"
)

// func to handle product creation
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product entities.Product
	json.NewDecoder(r.Body).Decode(&product)
	database.Instance.Create(&product)
	// encode response as json , i guess
	json.NewEncoder(w).Encode(product)

}
