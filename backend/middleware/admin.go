package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AdminMiddleware(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	tipo := claims["type"].(string)

	if tipo != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "No tienes permiso para realizar esta acción.",
		})
	}

	return c.Next()
}
