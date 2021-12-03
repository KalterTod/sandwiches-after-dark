package menu

import (
    "encoding/json"
    "net/http"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "context"
    "sandwiches-after-dark/db"
    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

var Client *mongo.Client
var CTX context.Context

func Menu(response http.ResponseWriter, request *http.Request) {
    // Set up content-type and define collection
    response.Header().Set("content-type", "application/json")
    collection := db.Session.Database("sandwiches-after-dark").Collection("menu")

    // Create cursor to be used for result-set and set up error handling
    cursor, currErr := collection.Find(CTX, bson.M{})

    if currErr != nil {
        response.WriteHeader(http.StatusInternalServerError)
        response.Write([]byte(`{"message": "` + currErr.Error() + `"}`))
    }
    defer cursor.Close(CTX)

    var items []MenuItem
    for cursor.Next(CTX) {
        var menuItem MenuItem
        cursor.Decode(&menuItem)
        items = append(items, menuItem)
    }

    if err := cursor.Err(); err != nil {
        response.WriteHeader(http.StatusInternalServerError)
        response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
        return
    }

    json.NewEncoder(response).Encode(items)
}

func Item(response http.ResponseWriter, request *http.Request) {
  response.Header().Set("content-type", "application/json")
  params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
  collection := db.Session.Database("sandwiches-after-dark").Collection("menu")

  var menuItem MenuItem
  err := collection.FindOne(CTX, MenuItem{ID: id}).Decode(&menuItem)

  if err != nil {
      response.WriteHeader(http.StatusInternalServerError)
      response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
      return
  }

  json.NewEncoder(response).Encode(menuItem)
}
