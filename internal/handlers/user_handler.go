package handlers

import (
	"api.droppy.com.br/internal/models"
	"api.droppy.com.br/internal/repositories"
	"strconv"

	"api.droppy.com.br/internal/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userRepository *repositories.UserRepository
}

func NewUserHandler(userRepository *repositories.UserRepository) *UserHandler {
	return &UserHandler{userRepository: userRepository}
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userRepository.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao buscar usuários"})
	}

	if users == nil {
		users = []models.User{}
	}

	return c.JSON(fiber.Map{
		"data":  users,
		"total": len(users),
	})
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
	}

	user, err := h.userRepository.GetUserByID(id)
	if err != nil || user == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Usuário não encontrado"})
	}

	return c.JSON(user)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user services.UserService
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Dados inválidos"})
	}

	if err := h.userRepository.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao criar usuário"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Usuário criado com sucesso"})
}
