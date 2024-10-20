package routes

import (
	"todo-service/controllers"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(app *fiber.App, mongoClient *mongo.Client) {
	// Todo routes
	setupTodoRoutes(app)
}
func setupTodoRoutes(app *fiber.App) {
	todos := app.Group("/todos")
	todos.Get("/", controllers.GetTodos)
	todos.Get("/:id", controllers.GetTodoByID)
	todos.Post("/", controllers.CreateTodo)
	todos.Put("/:id", controllers.UpdateTodo)
	todos.Delete("/:id", controllers.DeleteTodo)
}
