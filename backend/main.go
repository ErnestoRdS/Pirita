// Pirita Backend es un binario hecho en Go que sirve como backend para la
// plataforma de Pirita. Este binario se encarga de servir las rutas protegidas
// de la API, así como las rutas públicas de la misma. Maneja la autenticación
// y autorización de los usuarios, así como la creación de los mismos.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	flogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	"github.com/UpVent/Pirita/v2/models"
	"github.com/UpVent/Pirita/v2/routes"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// main es la función principal de pirita y es el punto de entrada del programa.
// Desde aquí servimos las rutas protegidas, así como las rutas públicas de la
// API.
func main() {

	// Banderas de linea de comandos.
	var (
		// Bandera para cambiar el puerto de escucha.
		port string
	)

	flag.StringVar(&port, "port", "8080", "Cambia el puerto de escucha del programa.")
	flag.Usage = func() {
		fmt.Printf("Uso: %s [opciones]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	// Localización de la base de datos.
	dsn := "./db/db.sqlite"

	// Abrir la base de datos.
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	// Si hay un error al conectarse a la base de datos, terminar el programa.
	if err != nil {
		log.Fatalf("Error al conectarse a la base de datos: %v", err)
		os.Exit(1)
	}

	// Migrar la base de datos.
	err = db.AutoMigrate(
		&models.Conductor{},
		&models.Contrato{},
		&models.Pago{},
		&models.Vehiculo{},
		&models.Viaje{},
	)

	if err != nil {
		log.Fatalf("Error al migrar los modelos a la base de datos: %v", err)
	}

	// Crear la aplicación de Fiber.
	app := fiber.New()

	// Middelewares.
	// Usar el middleware de logger y el de limitador de peticiones.
	app.Use(flogger.New())
	app.Use(limiter.New())

	// Montar las rutas.
	routes.ConductorRouter(app, db)
	routes.ContratoRouter(app, db)
	routes.PagosRouter(app, db)
	routes.VehiculoRouter(app, db)
	routes.ViajeRouter(app, db)

	// Ruta de monitoreo.
	app.Get("/monitor", monitor.New(monitor.Config{
		Title: "Pirita Backend - Monitoreo",
	}))

	// Mostrar siempre un 404 en la ruta raíz. Esto es solo una AP que
	// recibe y responde JSON, no hay necesidad de mostrar una página
	// de inicio o algo por el estilo.
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	// Escuchar en el puerto especificado. (Por defecto 8080)
	log.Fatal(app.Listen(port))
}
