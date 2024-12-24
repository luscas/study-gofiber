package routes

import (
	"api.droppy.com.br/internal/handlers"
	"api.droppy.com.br/internal/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	userRepository := repositories.NewUserRepository(db)

	userHandler := handlers.NewUserHandler(userRepository)

	SetupUserRoutes(app, userHandler)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
}
