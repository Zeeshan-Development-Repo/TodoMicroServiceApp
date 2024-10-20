package controllers

import (
	"todo-service/jwt_service"
	"todo-service/models"
	"todo-service/services"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TodoController is the structure that holds the TodoService
type TodoController struct {
	service services.TodoService
}

// NewTodoController creates a new TodoController
func NewTodoController(service services.TodoService) *TodoController {
	return &TodoController{service: service}
}

// CreateTodoHandler handles the creation of a new todo item
func (c *TodoController) CreateTodoHandler(ctx *fiber.Ctx) error {
	var todo models.Todo
	if err := ctx.BodyParser(&todo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	createdTodo, err := c.service.CreateTodo(&todo)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create todo"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(createdTodo)
}

// GetTodoHandler handles retrieving a todo by ID
func (c *TodoController) GetTodoHandler(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	todo, err := c.service.GetTodo(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	}

	return ctx.JSON(todo)
}

// UpdateTodoHandler handles updating an existing todo item
func (c *TodoController) UpdateTodoHandler(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var updatedTodo models.Todo
	if err := ctx.BodyParser(&updatedTodo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	todo, err := c.service.UpdateTodo(id, updatedTodo)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update todo"})
	}

	return ctx.JSON(todo)
}

// DeleteTodoHandler handles deleting a todo item
func (c *TodoController) DeleteTodoHandler(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	err = c.service.DeleteTodo(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	}

	return ctx.Status(fiber.StatusNoContent).SendString("")
}

// GetTodosHandler retrieves all todos for a specific user
func (c *TodoController) GetTodosHandler(ctx *fiber.Ctx) error {
	claims := ctx.Locals("userClaims").(jwt.MapClaims)
	user, err := jwt_service.ExtractUserFromClaims(claims)

	if err != nil {
		println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing user_id parameter"})
	}

	if user.Id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing user_id parameter"})
	}

	todos, err := c.service.GetTodos(user.Id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve todos"})
	}

	return ctx.JSON(todos)
}
