package routes

import (
	"todo-service/controllers"
	"todo-service/jwt_service"
	"todo-service/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Setup Protection
	app.Use(jwt_service.ProtectedRoute())

	// Todo routes
	setupTodoRoutes(app)
}
func setupTodoRoutes(app *fiber.App) {
	todos := app.Group("/todos")

	// Create a single instance of the TodoService and TodoController
	todoService := services.NewTodoService()
	todoController := controllers.NewTodoController(todoService)
	println("Dependencies Created!")

	todos.Get("/", todoController.GetTodosHandler)
	todos.Get("/:id", todoController.GetTodoHandler)
	todos.Post("/", todoController.CreateTodoHandler)
	todos.Put("/:id", todoController.UpdateTodoHandler)
	todos.Delete("/:id", todoController.DeleteTodoHandler)
}
