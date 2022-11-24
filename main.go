package main

import (
	"fmt"
	"golang-crud-rest-api/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// this repo intended as learning with ATM method from
// https://codewithmukesh.com/blog/implementing-crud-in-golang-rest-api/
func main() {
	LoadAppConfig()
	database.Connect(appConf.Connection)
	database.Migrate()

	// initialize router
	router := mux.NewRouter().StrictSlash(true)

	// Register router

	RegisterProductRoute(router)

	// start the server

	log.Println(fmt.Sprintf("starting server on the port %s", appConf.Port))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", appConf.Port), router))

}

func RegisterProductRoute(router *mux.Router) {

	router.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", controllers.GetProductsById).Methods("GET")
	router.HandleFunc("/api/products", controllers.GetProducts).Methods("POST")
	router.HandleFunc("/api/products/{id}", controllers.GetProducts).Methods("PUT")
	router.HandleFunc("/api/products/{id}", controllers.GetProducts).Methods("DELETE")
}
