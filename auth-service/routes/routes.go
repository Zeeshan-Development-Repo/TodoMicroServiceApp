package routes

import (
	"auth-service/controller"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes initializes all routes for the app
func SetupRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	setupAuthRoutes(auth)
}

func setupAuthRoutes(auth fiber.Router) {
	auth.Post("/login/basic", controller.BasicAuthLogin)
	auth.Post("/signup/basic", controller.BasicCreateAccount)
	auth.Post("/verify/basic", controller.BasicVerifyToken)
}
