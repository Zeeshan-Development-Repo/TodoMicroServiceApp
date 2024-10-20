package controllers

import (
	"user-service/models"

	"github.com/gofiber/fiber/v2"
)

// Get all users
func GetUsers(c *fiber.Ctx) error {
	users := models.GetAllUsers()
	return c.Status(fiber.StatusOK).JSON(users)
}

// Get user by ID
func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := models.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

// Create new user
func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	models.AddUser(user)
	return c.Status(fiber.StatusCreated).JSON(user)
}

// Update user
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	updatedUser, err := models.UpdateUser(id, user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}
	return c.Status(fiber.StatusOK).JSON(updatedUser)
}

// Delete user
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	err := models.DeleteUser(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User deleted"})
}
