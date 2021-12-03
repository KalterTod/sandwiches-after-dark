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

func Menu(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("content-type", "application/json")
    collection := db.Session.Database("sandwiches-after-dark").Collection("menu")
    cursor, currErr := collection.Find(context.TODO(), bson.M{})

    if currErr != nil {
        response.WriteHeader(http.StatusInternalServerError)
        response.Write([]byte(`{"message": "` + currErr.Error() + `"}`))
    }
    defer cursor.Close(context.TODO())

    var items []MenuItem
    for cursor.Next(context.TODO()) {
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
  err := collection.FindOne(context.TODO(), MenuItem{ID: id}).Decode(&menuItem)

  if err != nil {
      response.WriteHeader(http.StatusInternalServerError)
      response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
      return
  }

  json.NewEncoder(response).Encode(menuItem)
}
