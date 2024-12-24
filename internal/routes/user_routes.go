package routes

import (
	"api.droppy.com.br/internal/handlers"
	"api.droppy.com.br/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app fiber.Router, userHandler *handlers.UserHandler) {
	users := app.Group("/users")

	users.Use(middlewares.AuthMiddleware)

	users.Get("/", userHandler.GetUsers)
	users.Get("/:id", userHandler.GetUser)
	users.Post("/", userHandler.CreateUser)
}
