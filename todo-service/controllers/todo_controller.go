package controllers

import (
	"todo-service/models"

	"github.com/gofiber/fiber/v2"
)

// Get all todos
func GetTodos(c *fiber.Ctx) error {
	todos := models.GetAllTodos()
	return c.Status(fiber.StatusOK).JSON(todos)
}

// Get todo by ID
func GetTodoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	todo, err := models.GetTodoByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Todo not found"})
	}
	return c.Status(fiber.StatusOK).JSON(todo)
}

// Create new todo
func CreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)

	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	models.AddTodo(todo)
	return c.Status(fiber.StatusCreated).JSON(todo)
}

// Update todo
func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	todo := new(models.Todo)

	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	updatedTodo, err := models.UpdateTodo(id, todo)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Todo not found"})
	}
	return c.Status(fiber.StatusOK).JSON(updatedTodo)
}

// Delete todo
func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	err := models.DeleteTodo(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Todo not found"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Todo deleted"})
}
