package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")

	auth.Post("/register", handlers.RegisterUser)
	auth.Post("/login", handlers.LoginUser)
}
