package routes

import (
	"auth-service/controller"
	"auth-service/jwt_utils"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes initializes all routes for the app
func SetupRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	setupGoogleAuthRoutes(auth)
	setupFacebookAuthRoutes(auth)
	setupAuthRoutes(auth)
}

func setupAuthRoutes(auth fiber.Router) {
	auth.Post("/login/basic", controller.BasicAuthLogin)
	auth.Post("/signup/basic", controller.BasicCreateAccount)
}

func setupGoogleAuthRoutes(auth fiber.Router) {
	auth.Get("/login/:provider", controller.GoogleLogin)
	auth.Get("/callback/:provider", jwt_utils.HandleOAuthCallback)
}

func setupFacebookAuthRoutes(auth fiber.Router) {
	auth.Get("/login/:provider", controller.FacebookLogin)
	auth.Get("/callback/:provider", jwt_utils.HandleOAuthCallback)
}
