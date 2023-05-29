package routes

import (
	"github.com/UpVent/Pirita/v2/middleware"
	"github.com/UpVent/Pirita/v2/models"
	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

// PagosRouter es el router para la API de pagos de Pirita
//
// Las rutas de este router empiezan con `/api/pagos`
// y se pueden agrupar en:
// - Obtener todos los pagos (GET `/api/pagos`)
// - Obtener un pago por ID (GET `/api/pagos/:id`)
// - Crear un nuevo pago (POST `/api/pagos`)
// - Actualizar un pago por ID (PUT `/api/pagos/:id`)
// - Eliminar un pago por ID (DELETE `/api/pagos/:id`)
//
// Para más información sobre los contratos, ver el modelo `Pago`.
func PagosRouter(app *fiber.App, db *gorm.DB, jwtMiddleware fiber.Handler) {
	// Obtener todos los pagos.
	app.Get("/api/pagos", jwtMiddleware, middleware.AdminMiddleware, func(c *fiber.Ctx) error {
		var pagos []models.Pago
		if err := db.Find(&pagos).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error al obtener los pagos",
			})
		}
		return c.JSON(pagos)
	})

	// Obtener un pago por ID.
	app.Get("/api/pagos/:id", jwtMiddleware, func(c *fiber.Ctx) error {
		var pago models.Pago
		if err := db.Find(&pago, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Pago no encontrado",
			})
		}
		return c.JSON(pago)
	})

	// Crear un nuevo pago.
	app.Post("/api/pagos", jwtMiddleware, middleware.AdminMiddleware, func(c *fiber.Ctx) error {
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
	app.Put("/api/pagos/:id", jwtMiddleware, middleware.AdminMiddleware, func(c *fiber.Ctx) error {
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
	app.Delete("/pagos/:id", jwtMiddleware, middleware.AdminMiddleware, func(c *fiber.Ctx) error {
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
