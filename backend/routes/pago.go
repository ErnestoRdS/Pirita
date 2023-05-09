package routes

import (
	"github.com/UpVent/Pirita/v2/models"
	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

func PagosRouter(app *fiber.App, db *gorm.DB) {
	// Obtener todos los pagos.
	app.Get("/api/pagos", func(c *fiber.Ctx) error {
		var pagos []models.Pago
		if err := db.Find(&pagos).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error al obtener los pagos",
			})
		}
		return c.JSON(pagos)
	})

	// Obtener un pago por ID.
	app.Get("/api/pagos/:id", func(c *fiber.Ctx) error {
		var pago models.Pago
		if err := db.Find(&pago, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Pago no encontrado",
			})
		}
		return c.JSON(pago)
	})

	// Crear un nuevo pago.
	app.Post("/api/pagos", func(c *fiber.Ctx) error {
		var pago models.Pago

		if err := c.BodyParser(&pago); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Error al analizar el cuerpo de la solicitud",
			})
		}
		if err := db.Create(&pago).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error al crear el pago",
			})
		}
		return c.JSON(pago)
	})

	// Actualizar un pago por ID.
	app.Put("/api/pagos/:id", func(c *fiber.Ctx) error {
		var pago models.Pago
		if err := db.First(&pago, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Pago no encontrado",
			})
		}
		if err := c.BodyParser(&pago); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Error al analizar el cuerpo de la solicitud",
			})
		}
		if err := db.Save(&pago).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error al actualizar el pago",
			})
		}
		return c.JSON(pago)
	})

	// Eliminar un pago por ID
	app.Delete("/pagos/:id", func(c *fiber.Ctx) error {
		var pago models.Pago
		if err := db.First(&pago, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Pago no encontrado",
			})
		}
		if err := db.Delete(&pago).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error al eliminar el pago",
			})
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
