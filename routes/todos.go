package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theascendedeber/TodoListAPI/middlewares"
)

func RegisterTodoRoutes(app *fiber.App) {
	todos := app.Group("/todos")

	todos.Use(middlewares.RequireLogin)

	todos.Get("/", handlers.GetTodos)
	todos.Post("/", handlers.CreateTodo)
	todos.Put("/:id", hanlers.UpdateTodo)
	todos.Delete("/:id", handlers.DeleteTodo)
	todos.Get("/:id", handlers.GetTodoByID)
}
