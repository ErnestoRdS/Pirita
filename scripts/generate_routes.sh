#!/bin/bash

# Obtener la ruta de los modelos y las rutas del backend
MODELS_PATH="./backend/models"
ROUTES_PATH="./backend/routes"

# Si la cantidad de archivos en ambos directorios es la misma, entonces no hay que hacer nada.
if [ "$(ls -1q $MODELS_PATH | wc -l)" -eq "$(ls -1q $ROUTES_PATH | wc -l)" ]; then
	echo "Ya existe una cantidad igual de modelos y rutas."
	exit 0
fi

for MODEL in "$MODELS_PATH"/*.go; do
	MODEL_NAME=$(basename "$MODEL" .go)

	ROUTES_FILE="$ROUTES_PATH/$MODEL_NAME.go"

    ROUTES_CONTENT=$(cat <<-EOF
	package routes

	import (
		   "github.com/gofiber/fiber/v2"
		   "github.com/UpVent/Pirita/v2/models"

		   "gorm.io/gorm"
	)

	func ${MODEL_NAME}Router(app *fiber.App, db *gorm.DB) {
		 // Obtener todos los ${MODEL_NAME}s
		 app.Get("/api/${MODEL_NAME}s", func(c *fiber.Ctx) error {
		    var ${MODEL_NAME}s []models.${MODEL_NAME}
			if err := db.Find(&${MODEL_NAME}s).Error; err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error al obtener los ${MODEL_NAME}s",
				})
			}
				return c.JSON(${MODEL_NAME}s)
		})

		// Obtener un ${MODEL_NAME} por ID
		app.Get("/api/${MODEL_NAME}s/:id", func(c *fiber.Ctx) error {
			id := c.Params("id")
			var ${MODEL_NAME} models.${MODEL_NAME}
			if err := db.First(&${MODEL_NAME}, id).Error; err != nil {
			   return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			   	"message": "No se encontró el ${MODEL_NAME} con ID " + id,
					})
				}
			return c.JSON(${MODEL_NAME})
		})

		// Crear un nuevo ${MODEL_NAME}
		app.Post("/api/${MODEL_NAME}s", func(c *fiber.Ctx) error {
			var ${MODEL_NAME} models.${MODEL_NAME}
			if err := c.BodyParser(&${MODEL_NAME}); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Error al decodificar el cuerpo de la petición",
				})
			}
			if err := db.Create(&${MODEL_NAME}).Error; err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						message": "Error al crear el ${MODEL_NAME}",
				})
			}
			return c.JSON(${MODEL_NAME})
		})

		// Actualizar un ${MODEL_NAME} por ID
		app.Put("/api/${MODEL_NAME}s/:id", func(c *fiber.Ctx) error {
			id := c.Params("id")
			var ${MODEL_NAME} models.${MODEL_NAME}
			if err := db.First(&${MODEL_NAME}, id).Error; err != nil {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"message": "No se encontró el ${MODEL_NAME} con ID " + id,
				})
			}
			if err := c.BodyParser(&${MODEL_NAME}); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Error al decodificar el cuerpo de la petición",
				})
			}
			if err := db.Save(&${MODEL_NAME}).Error; err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error al actualizar el ${MODEL_NAME}",
				})
			}
			return c.JSON(${MODEL_NAME})
		})

		// Eliminar un ${MODEL_NAME} por ID
		app.Delete("/api/${MODEL_NAME}s/:id", func(c *fiber.Ctx) error {
			id := c.Params("id")
			var ${MODEL_NAME} models.${MODEL_NAME}
			if err := db.First(&${MODEL_NAME}, id).Error; err != nil {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"message": "No se encontró el ${MODEL_NAME} con ID " + id,
				})
			}
			if err := db.Delete(&${MODEL_NAME}).Error; err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error al eliminar el ${MODEL_NAME}",
				})
			}
			return c.SendStatus(fiber.StatusNoContent)
		})
	}
EOF
)

	# Si ya existe el archivo, entonces no hay que crearlo.
	if [ -f "$ROUTES_FILE" ]; then
		echo "El archivo $ROUTES_FILE ya existe."
		continue
	else
		echo "Creando archivo $ROUTES_FILE..."
		echo "$ROUTES_CONTENT" > "$ROUTES_FILE"
	fi

	echo "Formateando código generado..."
	gofmt -w ./backend/.
done
