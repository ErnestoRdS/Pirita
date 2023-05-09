package routes

import (
	"github.com/UpVent/Pirita/v2/models"
	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

func contratoRouter(app *fiber.App, db *gorm.DB) {
	// Obtener todos los contratos
	app.Get("/api/contratos", func(c *fiber.Ctx) error {
		var contratos []models.Contrato
		if err := db.Find(&contratos).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error al obtener los contratos",
			})
		}
		return c.JSON(contratos)
	})

	// Obtener un contrato por ID
	app.Get("/api/contratos/:id", func(c *fiber.Ctx) error {
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
	app.Post("/api/contratos", func(c *fiber.Ctx) error {
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
	app.Put("/api/contratos/:id", func(c *fiber.Ctx) error {
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
	app.Delete("/api/contratos/:id", func(c *fiber.Ctx) error {
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
