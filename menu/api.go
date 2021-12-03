package menu

import (
  "fmt"
  "net/http"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/bson"
  "context"
  "sandwiches-after-dark/db"
)

var Client *mongo.Client

func Menu(w http.ResponseWriter, r *http.Request) {
  collection := db.Session.Database("sandwiches-after-dark").Collection("menu")
  cur, currErr := collection.Find(context.TODO(), bson.M{})

  if currErr != nil { panic(currErr) }
  defer cur.Close(context.TODO())

  var items []MenuItem
  cur.All(context.TODO(), &items)
  fmt.Println(items)
}

func ItemOptions(w http.ResponseWriter, r *http.Request) {
  // TODO: Implement menu item options endpoint
}
