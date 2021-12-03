package db

import (
  "fmt"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "context"
  "log"
  "os"
)

var Session = Connection()

func Connection() *mongo.Client {
    // Set client options
    clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URL"))

    // Connect to MongoDB
    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil {
      log.Fatal(err)
    }

    // Check the connection
    err = client.Ping(context.TODO(), nil)

    if err != nil {
      log.Fatal(err)
    }

    // defer client.Disconnect(context.TODO())

    fmt.Println("Connected to MongoDB!")

    return client
  }
