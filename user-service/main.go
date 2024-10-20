package main

import (
	"context"
	"log"
	"os"
	"user-service/routes"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// MongoDB connection
	mongoURI := os.Getenv("MONGODB_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Set up routes
	routes.SetupRoutes(app, client)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
