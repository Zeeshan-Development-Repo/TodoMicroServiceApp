package controller

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shareed2k/goth_fiber"
)

func GoogleLogin(c *fiber.Ctx) error {
	goth_fiber.BeginAuthHandler(c)
	return nil
}
