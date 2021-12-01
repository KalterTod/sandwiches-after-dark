package main

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
)

type Thing struct {
  Title string `json:"Title"`
}

type StatusMessage struct {
  Msg string `json:"Msg"`
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
  healthcheck_message := StatusMessage{Msg: "OK"}
  json.NewEncoder(w).Encode(healthcheck_message)
}

func getThings(w http.ResponseWriter, r *http.Request) {
  return_thing := Thing{Title: "Hello World"}
  json.NewEncoder(w).Encode(return_thing)
}

func handleRequests() {
  fmt.Println("Hello World")
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/get-things", getThings)
  router.HandleFunc("/healthcheck", healthcheck)
  log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
  fmt.Println("Rest API v2.0 - Mux Routers")
  handleRequests()
}
