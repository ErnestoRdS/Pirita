package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/UpVent/Pirita/v2/models"
	"github.com/UpVent/Pirita/v2/routes"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Connectarse a la base de datos.
	dsn := "./db/db.sqlite"

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Error al conectarse a la base de datos: %v", err)
	}

	// Migrar la base de datos.
	err = db.AutoMigrate(
		&models.Conductor{},
		&models.Pago{},
	)

	if err != nil {
		log.Fatalf("Error al migrar los modelos a la base de datos: %v", err)
	}

	app := fiber.New()

	// Montar las rutas.
	routes.PagosRouter(app, db)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	app.Listen(":3000")
}
