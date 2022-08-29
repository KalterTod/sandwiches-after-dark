package user

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

var Client *mongo.Client
var CTX context.Context

func User(response http.ResponseWriter, request *http.Request) {
	//TODO: Retrieve user and favorite orders
}

func CreateUser(response http.ResponseWriter, request *http.Request) {
	//TODO: User Create
}

func EditUser(response http.ResponseWriter, request *http.Request) {
	//TODO: Give user ability to update personal information
}

func Login(response http.ResponseWriter, request *http.Request) {
	//TODO: Add login through LDAP
}
