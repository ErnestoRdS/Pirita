package routes

import (
	"log"

	"gorm.io/gorm"
	"github.com/gofiber/fiber/v2"

)

// ObservabilityRouter es el router para la API de observabilidad de Pirita.
// Las rutas de este router empiezan con `/api/diagnose/` y se pueden agrupar en:
// - TODO: Listar más endpoints.
func ObservabilityRouter(app *fiber.App, db *gorm.DB) {
	app.Get("/api/diagnose/ping", func(c *fiber.Ctx) error {
		log.Printf("[INFO] [%s] - ¡Se hizo un ping a la API!\n", c.IP())
		return c.SendString("pong")
	})
}
