package routes

import (
	"github.com/UpVent/Pirita/v2/middleware"
	"github.com/UpVent/Pirita/v2/models"
	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

// ViajeRouter es el router para la API de viajes de Pirita
//
// Las rutas de este router empiezan con `/api/viajes`
// y se pueden agrupar en:
// - Obtener todos los viajes (GET `/api/viajes`)
// - Obtener un viaje por ID (GET `/api/viajes/:id`)
// - Crear un nuevo viaje (POST `/api/viajes`)
// - Actualizar un viaje por ID (PUT `/api/viajes/:id`)
// - Eliminar un viaje por ID (DELETE `/api/viajes/:id`)
//
// Para más información sobre los contratos, ver el modelo `Viaje`.
func ViajeRouter(app *fiber.App, db *gorm.DB, jwtMiddleware fiber.Handler) {
	// Obtener todos los viajes
	app.Get("/api/viajes", jwtMiddleware, middleware.AdminMiddleware, func(c *fiber.Ctx) error {
		var viajes []models.Viaje
		if err := db.Find(&viajes).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error al obtener los viajes",
			})
		}
		return c.JSON(viajes)
	})

	// Obtener un viaje por ID
	app.Get("/api/viajes/:id", jwtMiddleware, func(c *fiber.Ctx) error {
		id := c.Params("id")
		var viaje models.Viaje
		if err := db.First(&viaje, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "No se encontró el viaje con ID " + id,
			})
		}
		return c.JSON(viaje)
	})

	// Crear un nuevo viaje
	app.Post("/api/viajes", jwtMiddleware, middleware.AdminMiddleware, func(c *fiber.Ctx) error {
		var viaje models.Viaje
		if err := c.BodyParser(&viaje); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Error al decodificar el cuerpo de la petición",
			})
		}
		if err := db.Create(&viaje).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error al crear el viaje",
			})
		}
		return c.JSON(viaje)
	})

	// Actualizar un viaje por ID
	app.Put("/api/viajes/:id", jwtMiddleware, middleware.AdminMiddleware, func(c *fiber.Ctx) error {
		id := c.Params("id")
		var viaje models.Viaje
		if err := db.First(&viaje, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "No se encontró el viaje con ID " + id,
			})
		}
		if err := c.BodyParser(&viaje); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Error al decodificar el cuerpo de la petición",
			})
		}
		if err := db.Save(&viaje).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error al actualizar el viaje",
			})
		}
		return c.JSON(viaje)
	})

	// Eliminar un viaje por ID
	app.Delete("/api/viajes/:id", jwtMiddleware, middleware.AdminMiddleware, func(c *fiber.Ctx) error {
		id := c.Params("id")
		var viaje models.Viaje
		if err := db.First(&viaje, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "No se encontró el viaje con ID " + id,
			})
		}
		if err := db.Delete(&viaje).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error al eliminar el viaje",
			})
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
