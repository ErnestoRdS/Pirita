package routes

import (
	"github.com/UpVent/Pirita/v2/models"
	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

func vehiculoRouter(app *fiber.App, db *gorm.DB) {
	// Obtener todos los vehiculos
	app.Get("/api/vehiculos", func(c *fiber.Ctx) error {
		var vehiculos []models.Vehiculo
		if err := db.Find(&vehiculos).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error al obtener los vehiculos",
			})
		}
		return c.JSON(vehiculos)
	})

	// Obtener un vehiculo por ID
	app.Get("/api/vehiculos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var vehiculo models.Vehiculo
		if err := db.First(&vehiculo, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "No se encontró el vehiculo con ID " + id,
			})
		}
		return c.JSON(vehiculo)
	})

	// Crear un nuevo vehiculo
	app.Post("/api/vehiculos", func(c *fiber.Ctx) error {
		var vehiculo models.Vehiculo
		if err := c.BodyParser(&vehiculo); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Error al decodificar el cuerpo de la petición",
			})
		}
		if err := db.Create(&vehiculo).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error al crear el vehiculo",
			})
		}
		return c.JSON(vehiculo)
	})

	// Actualizar un vehiculo por ID
	app.Put("/api/vehiculos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var vehiculo models.Vehiculo
		if err := db.First(&vehiculo, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "No se encontró el vehiculo con ID " + id,
			})
		}
		if err := c.BodyParser(&vehiculo); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Error al decodificar el cuerpo de la petición",
			})
		}
		if err := db.Save(&vehiculo).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error al actualizar el vehiculo",
			})
		}
		return c.JSON(vehiculo)
	})

	// Eliminar un vehiculo por ID
	app.Delete("/api/vehiculos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var vehiculo models.Vehiculo
		if err := db.First(&vehiculo, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "No se encontró el vehiculo con ID " + id,
			})
		}
		if err := db.Delete(&vehiculo).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error al eliminar el vehiculo",
			})
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
