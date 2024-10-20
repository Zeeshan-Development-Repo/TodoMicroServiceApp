package config

import (
	"auth-service/models"
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Database() {
	// mongoURI := "mongodb://localhost:27017"
	mongoURI := os.Getenv("MONGODB_URI")

	var err error

	// Create a new MongoDB client and connect to the database
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	// Set the user collection
	models.UserCollection = Client.Database("auth-db").Collection("users")
}

// Disconnect the MongoDB client
func DisconnectDatabase() {
	if err := Client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
}
