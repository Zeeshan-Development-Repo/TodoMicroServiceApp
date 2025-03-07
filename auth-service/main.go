package main

import (
	"auth-service/config"
	"auth-service/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.Database()
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3002"))
}
