package middlewares

import (
	"encoding/json"
	"strings"

	"api.droppy.com.br/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		errorMessage, _ := json.Marshal(fiber.Map{
			"error": "Missing Authorization header",
		})
		return fiber.NewError(fiber.StatusUnauthorized, string(errorMessage))
	}

	bearerToken := strings.Split(authHeader, "Bearer ")
	if len(bearerToken) != 2 {
		errorMessage, _ := json.Marshal(fiber.Map{
			"error": "Missing or invalid Bearer token",
		})
		return fiber.NewError(fiber.StatusUnauthorized, string(errorMessage))
	}

	token := bearerToken[1]
	err := jwt.VerifyToken(token)
	if err != nil {
		errorMessage, _ := json.Marshal(fiber.Map{
			"error": "Unauthorized",
		})
		return fiber.NewError(fiber.StatusUnauthorized, string(errorMessage))
	}

	return c.Next()
}
func Recover() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if err := recover(); err != nil {
				if fiberErr, ok := err.(*fiber.Error); ok {
					c.Status(fiberErr.Code).JSON(fiber.Map{"error": fiberErr.Message})
					return
				}
				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
			}
		}()
		return c.Next()
	}
}

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			return err
		}
		return nil
	}
}
