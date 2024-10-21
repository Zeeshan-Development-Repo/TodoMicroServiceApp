package config

import (
	"context"
	"log"
	"os"
	"todo-service/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Database() {
	// mongoURI := "mongodb://localhost:27017"
	// mongoURI := os.Getenv("MONGODB_URI")
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	var err error

	// Create a new MongoDB client and connect to the database
	Client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	// Set the Todo collection
	models.TodoCollection = Client.Database("auth-db").Collection("Todos")
}

// Disconnect the MongoDB client
func DisconnectDatabase() {
	if err := Client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
}
