package main

import (
	"log"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/theascendedeber/TodoListAPI/database"
	"github.com/theascendedeber/TodoListAPI/routes"
)

func main() {
	if err := database.Connect(); err != nil {
		slog.Error("Ошибка подключения к бд")
	}

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Use(logger.New())

	routes.RegisterAuthRoutes(app)
	// routes.RegisterTodoRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
