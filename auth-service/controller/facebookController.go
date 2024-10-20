package controller

import "github.com/gofiber/fiber/v2"

func FacebookLogin(c *fiber.Ctx) error {
	return c.SendString("Facebook login route")
}
