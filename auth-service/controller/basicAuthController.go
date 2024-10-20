package controller

import (
	"auth-service/jwt_service"
	"auth-service/services"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func BasicAuthLogin(c *fiber.Ctx) error {
	// Parse the request body to get the username and password
	var body LoginRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	// Validate the username and password (you can replace this logic with database checks)
	user, err := services.CheckUser(body.Username) // Assuming services.CheckUser checks in DB
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid credentials or user not found",
		})
	}

	// Compare password (assuming user.Password is stored in DB)
	pass_err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if pass_err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid credentials",
		})
	}

	token, token_err := jwt_service.CreateToken(
		user.ID.Hex(),
		user.Email,
		user.Name,
		user.Username,
	)

	if token_err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Login Failed!",
		})
	}

	// If the credentials are valid, return success
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Login successful",
		"token":   token,
	})
}

// Request body struct for creating an account
type CreateAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func BasicCreateAccount(c *fiber.Ctx) error {
	// Parse the request body to get user details
	var body CreateAccountRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	// Check if the username or email already exists
	_, err := services.CheckUser(body.Username)
	if err == nil { // If no error, user already exists
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  "error",
			"message": "User already exists",
		})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not hash password",
		})
	}

	err = services.CreateUser(body.Username, string(hashedPassword))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error creating user",
		})
	}

	// Return success response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "User created successfully",
	})
}

// Request body struct for creating an account
type BasicVerifyTokenRequest struct {
	Token string `json:"token"`
}

func BasicVerifyToken(c *fiber.Ctx) error {
	// Parse the request body to get user details
	var body BasicVerifyTokenRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	if _, err := jwt_service.VerifyToken(body.Token); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid Token Error from jwt service",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success",
		"message": "Token Verified",
	})
}
