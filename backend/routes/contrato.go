package routes

import (
	"github.com/UpVent/Pirita/v2/middleware"
	"github.com/UpVent/Pirita/v2/models"
	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

// ContratoRouter es el router para la API de contratos de Pirita
//
// Las rutas de este router empiezan con `/api/contratos`
// y se pueden agrupar en:
// - Obtener todos los contratos (GET `/api/contratos`)
// - Obtener un contrato por ID (GET `/api/contratos/:id`)
// - Crear un nuevo contrato (POST `/api/contratos`)
// - Actualizar un contrato por ID (PUT `/api/contratos/:id`)
// - Eliminar un contrato por ID (DELETE `/api/contratos/:id`)
//
// Para más información sobre los contratos, ver el modelo `Contrato`.
func ContratoRouter(app *fiber.App, db *gorm.DB, jwtMiddleware fiber.Handler) {
	// Obtener todos los contratos
	app.Get("/api/contratos", jwtMiddleware, middleware.AdminMiddleware, func(c *fiber.Ctx) error {
		var contratos []models.Contrato
		if err := db.Find(&contratos).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error al obtener los contratos",
			})
		}
		return c.JSON(contratos)
	})

	// Obtener un contrato por ID
	app.Get("/api/contratos/:id", jwtMiddleware, func(c *fiber.Ctx) error {
		id := c.Params("id")
		var contrato models.Contrato
		if err := db.First(&contrato, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "No se encontró el contrato con ID " + id,
			})
		}
		return c.JSON(contrato)
	})

	// Crear un nuevo contrato
	app.Post("/api/contratos", jwtMiddleware, middleware.AdminMiddleware, func(c *fiber.Ctx) error {
		var contrato models.Contrato
		if err := c.BodyParser(&contrato); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Error al decodificar el cuerpo de la petición",
			})
		}
		if err := db.Create(&contrato).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error al crear el contrato",
			})
		}
		return c.JSON(contrato)
	})

	// Actualizar un contrato por ID
	app.Put("/api/contratos/:id", jwtMiddleware, middleware.AdminMiddleware, func(c *fiber.Ctx) error {
		id := c.Params("id")
		var contrato models.Contrato
		if err := db.First(&contrato, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "No se encontró el contrato con ID " + id,
			})
		}
		if err := c.BodyParser(&contrato); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Error al decodificar el cuerpo de la petición",
			})
		}
		if err := db.Save(&contrato).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error al actualizar el contrato",
			})
		}
		return c.JSON(contrato)
	})

	// Eliminar un contrato por ID
	app.Delete("/api/contratos/:id", jwtMiddleware, middleware.AdminMiddleware, func(c *fiber.Ctx) error {
		id := c.Params("id")
		var contrato models.Contrato
		if err := db.First(&contrato, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "No se encontró el contrato con ID " + id,
			})
		}
		if err := db.Delete(&contrato).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error al eliminar el contrato",
			})
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
