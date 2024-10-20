package config

import (
	"context"
	"log"
	"todo-service/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Database() {
	mongoURI := "mongodb://localhost:27017"
	var err error

	// Create a new MongoDB client and connect to the database
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
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
