package main

import (
	"fmt"
	"log"
	"net/http"
	"sandwiches-after-dark/menu"
	"sandwiches-after-dark/user"

	"github.com/gorilla/mux"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/menu", menu.Menu).Methods("GET")
	router.HandleFunc("/item/{id}", menu.Item).Methods("GET")
	router.HandleFunc("/item", menu.PostItem).Methods("POST")
	router.HandleFunc("/item/{id}", menu.EditItem).Methods("PUT")

	router.HandleFunc("/user/{id}", user.User).Methods("GET")
	router.HandleFunc("/user", user.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", user.EditUser).Methods("PUT")
	router.HandleFunc("/login", user.Login).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	fmt.Println("Sandwiches After Dark API v1.0")
	handleRequests()
}
