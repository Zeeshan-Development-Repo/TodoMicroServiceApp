package main

import (
	"log"
	"todo-service/config"
	"todo-service/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	config.Database()

	// Set up routes
	routes.SetupRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
