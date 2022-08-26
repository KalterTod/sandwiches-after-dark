package main

import (
	"fmt"
	"log"
	"net/http"
	"sandwiches-after-dark/menu"

	"github.com/gorilla/mux"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/menu", menu.Menu).Methods("GET")
	router.HandleFunc("/item/{id}", menu.Item).Methods("GET")
	router.HandleFunc("/item", menu.PostItem).Methods("POST")
	router.HandleFunc("/item/{id}", menu.EditItem).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	fmt.Println("Sandwiches After Dark API v1.0")
	handleRequests()
}
