package main

import (
  "log"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
)

type Thing struct {
  Title string `json:"Title"`
}

func getThings(w http.ResponseWriter, r *http.Request) {
  return_thing := Thing{Title: "Hello World",}
  json.NewEncoder(w).Encode(return_thing)
}

func main() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/get-things", getThings).Methods("GET")
  log.Fatal(http.ListenAndServe(":8080", router))
}
