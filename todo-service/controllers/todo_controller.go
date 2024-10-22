package controllers

import (
	"todo-service/jwt_service"
	"todo-service/models"
	"todo-service/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoController struct {
	service services.TodoService
}

func NewTodoController(service services.TodoService) *TodoController {
	return &TodoController{service: service}
}

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// CreateTodoHandler handles the creation of a new todo item
func (c *TodoController) CreateTodoHandler(ctx *fiber.Ctx) error {
	var Body CreateTodoRequest
	if err := ctx.BodyParser(&Body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	user, err := jwt_service.GetUserFromClaims(ctx)
	if err != nil {
		return err
	}

	createdTodo, err := c.service.CreateTodo(Body.Title, Body.Description, user.Id)
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
	user, err := jwt_service.GetUserFromClaims(ctx)
	if err != nil {
		return err
	}
	todos, err := c.service.GetTodos(user.Id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve todos"})
	}
	return ctx.JSON(todos)
}
