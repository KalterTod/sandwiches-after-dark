package menu

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sandwiches-after-dark/db"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

	fmt.Println(menuItem)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(response).Encode(menuItem)
}

func PostItem(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var newItem MenuItem

	err := json.NewDecoder(request.Body).Decode(&newItem)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	collection := db.Session.Database("sandwiches-after-dark").Collection("menu")

	result, insertErr := collection.InsertOne(CTX, newItem)
	if insertErr != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(response).Encode(result)
}

func EditItem(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var newItem MenuItem
	err := json.NewDecoder(request.Body).Decode(&newItem)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	updateObject := bson.M{
		"Name":    newItem.Name,
		"Desc":    newItem.Desc,
		"Options": newItem.Options,
		"Active":  newItem.Active,
		"Image":   newItem.Image}

	collection := db.Session.Database("sandwiches-after-dark").Collection("menu")
	_, updateErr := collection.UpdateOne(CTX, MenuItem{ID: id}, bson.M{"$set": updateObject})
	if updateErr != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + updateErr.Error() + `" }`))
		return
	}

	var updatedItem MenuItem
	findErr := collection.FindOne(CTX, MenuItem{ID: id}).Decode(&updatedItem)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + findErr.Error() + `" }`))
		return
	}

	json.NewEncoder(response).Encode(updatedItem)
}
