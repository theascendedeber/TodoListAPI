package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theascendedeber/TodoListAPI/database"
	"github.com/theascendedeber/TodoListAPI/models"
	"github.com/theascendedeber/TodoListAPI/utils"
)

const (
	requestCreateUser = "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
)

func RegisterUser(c *fiber.Ctx) error {
	var id int
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	hashPassword, err := utils.HashPassword([]byte(user.Password))
	if err != nil {
		c.Status(500).SendString("Ошибка при генерации хеша пароля")
	}

	err = database.DB.QueryRow(requestCreateUser, user.Name, user.Email, hashPassword).Scan(&id)
	if err != nil {
		return c.Status(500).SendString("Ошибка вставки данных в БД")
	}

	token := utils.GenerateJWT(id, user.Name)

	return c.Status(201).JSON(fiber.Map{"token": token})
}
