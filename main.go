package main

import (
  	"fmt"
  	"github.com/gorilla/mux"
  	"log"
  	"net/http"
  	"sandwiches-after-dark/menu"
)

func handleRequests() {
  	router := mux.NewRouter().StrictSlash(true)
  	router.HandleFunc("/menu", menu.Menu).Methods("GET")
  	router.HandleFunc("/item/{id}", menu.Item).Methods("GET")
  	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
  	fmt.Println("Sandwiches After Dark API v1.0")
    handleRequests()
}
