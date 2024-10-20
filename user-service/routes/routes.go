package routes

import (
	"user-service/controllers"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(app *fiber.App, mongoClient *mongo.Client) {
	// User routes
	setupUserRoutes(app)
}

func setupUserRoutes(app *fiber.App) {
	users := app.Group("/users")
	users.Get("/", controllers.GetUsers)
	users.Get("/:id", controllers.GetUserByID)
	users.Post("/", controllers.CreateUser)
	users.Put("/:id", controllers.UpdateUser)
	users.Delete("/:id", controllers.DeleteUser)
}
