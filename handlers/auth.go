package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theascendedeber/TodoListAPI/database"
	"github.com/theascendedeber/TodoListAPI/models"
	"github.com/theascendedeber/TodoListAPI/utils"
)

const (
	requestCreateUser           = "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	requestGetCountUsersByEmail = "SELECT COUNT(*) FROM users WHERE email = $1"
	requestGetUserByEmail       = "SELECT id, password FROM users WHERE email = $1"
)

func RegisterUser(c *fiber.Ctx) error {
	var id, count int64
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	err := database.DB.QueryRow(requestGetCountUsersByEmail, user.Email).Scan(&count)
	if err != nil {
		return c.Status(500).SendString("Ошибка при проверке существования пользователя")
	}
	if count > 0 {
		return c.Status(409).SendString("Пользователь с таким Email уже существует")
	}

	hashPassword, err := utils.HashPassword([]byte(user.Password))
	if err != nil {
		return c.Status(500).SendString("Ошибка при генерации хеша пароля")
	}

	err = database.DB.QueryRow(requestCreateUser, user.Name, user.Email, hashPassword).Scan(&id)
	if err != nil {
		return c.Status(500).SendString("Ошибка вставки данных в БД")
	}

	token, err := utils.GenerateJWT(id)
	if err != nil {
		return c.Status(500).SendString("Ошибка при генерации токена")
	}

	return c.Status(201).JSON(fiber.Map{"token": token})
}

func LoginUser(c *fiber.Ctx) error {
	loginRequest := new(models.LoginRequest)
	var loginResponse models.LoginResponse

	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	err := database.DB.QueryRow(requestGetUserByEmail, loginRequest.Email).Scan(&loginResponse.ID, &loginResponse.Password)
	if err != nil {
		return c.Status(500).SendString("Ошибка при запросе к БД")
	}

	truePassword := utils.VerifyPassword([]byte(loginRequest.Password), []byte(loginResponse.Password))
	if !truePassword {
		return c.Status(401).SendString("Неверные авторизационные данные")
	}

	token, err := utils.GenerateJWT(loginResponse.ID)
	if err != nil {
		return c.Status(500).SendString("Ошибка при генерации токена")
	}

	return c.Status(200).JSON(fiber.Map{"token": token})
}
