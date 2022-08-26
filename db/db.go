package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Session = Connection()

func Connection() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URL"))
	CTX, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(CTX, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(CTX, nil)

	if err != nil {
		log.Fatal(err)
	}

	// defer client.Disconnect(context.TODO())
	fmt.Println("Connected to MongoDB!")

	return client
}
