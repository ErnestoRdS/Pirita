package routes

import (
	"github.com/UpVent/Pirita/v2/models"
	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

func conductorRouter(app *fiber.App, db *gorm.DB) {
	// Obtener todos los conductores
	app.Get("/api/conductors", func(c *fiber.Ctx) error {
		var conductors []models.Conductor
		if err := db.Find(&conductors).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error al obtener los conductors",
			})
		}
		return c.JSON(conductors)
	})

	// Obtener un conductor por ID
	app.Get("/api/conductors/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var conductor models.Conductor
		if err := db.First(&conductor, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "No se encontró el conductor con ID " + id,
			})
		}
		return c.JSON(conductor)
	})

	// Crear un nuevo conductor
	app.Post("/api/conductors", func(c *fiber.Ctx) error {
		var conductor models.Conductor
		if err := c.BodyParser(&conductor); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Error al decodificar el cuerpo de la petición",
			})
		}
		if err := db.Create(&conductor).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error al crear el conductor",
			})
		}
		return c.JSON(conductor)
	})

	// Actualizar un conductor por ID
	app.Put("/api/conductors/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var conductor models.Conductor
		if err := db.First(&conductor, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "No se encontró el conductor con ID " + id,
			})
		}
		if err := c.BodyParser(&conductor); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Error al decodificar el cuerpo de la petición",
			})
		}
		if err := db.Save(&conductor).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error al actualizar el conductor",
			})
		}
		return c.JSON(conductor)
	})

	// Eliminar un conductor por ID
	app.Delete("/api/conductors/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var conductor models.Conductor
		if err := db.First(&conductor, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "No se encontró el conductor con ID " + id,
			})
		}
		if err := db.Delete(&conductor).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error al eliminar el conductor",
			})
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}