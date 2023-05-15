// Pirita Backend es un binario hecho en Go que sirve como backend para la
// plataforma de Pirita. Este binario se encarga de servir las rutas protegidas
// de la API, así como las rutas públicas de la misma. Maneja la autenticación
// y autorización de los usuarios, así como la creación de los mismos.
package main

import (
	"log"
	"net"
	"os"
	"runtime"

	"github.com/gofiber/fiber/v2"

	"github.com/UpVent/Pirita/v2/models"
	"github.com/UpVent/Pirita/v2/routes"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// main es la función principal de pirita y es el punto de entrada del programa.
// Desde aquí servimos las rutas protegidas, así como las rutas públicas de la
// API. Todo desde un UNIX socket.
func main() {
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

	app := fiber.New()

	// Montar las rutas.
	routes.ConductorRouter(app, db)
	routes.ContratoRouter(app, db)
	routes.PagosRouter(app, db)
	routes.VehiculoRouter(app, db)
	routes.ViajeRouter(app, db)

	// Rutas de observabilidad.
	routes.ObservabilityRouter(app, db)

	// Mostrar siempre un 404 en la ruta raíz.
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	var listener net.Listener

	switch runtime.GOOS {
	case "linux":
		listener, err = net.Listen("unix", "/tmp/pirita.sock")

		if err != nil {
			log.Fatalf("Error al crear el socket UNIX: %v", err)
			os.Exit(1)
		}

		app.Listener(listener)

		log.Printf("Iniciado en el socket UNIX: %v, sistema operativo: %v",
			listener.Addr(),
			runtime.GOOS)

		defer listener.Close()

	case "windows":
		listener, err = net.Listen("tcp", "8080")

		if err != nil {
			log.Fatalf("Error al crear el socket TCP: %v", err)
			os.Exit(1)
		}

		log.Printf("Iniciado en el socket TCP: %v, sistema operativo: %v",
			listener.Addr(),
			runtime.GOOS)

		app.Listener(listener)

		defer listener.Close()
	case "darwin":
		listener, err = net.Listen("unix", "/tmp/pirita.sock")

		if err != nil {
			log.Fatalf("Error al crear el socket UNIX: %v", err)
			os.Exit(1)
		}

		app.Listener(listener)

		log.Printf("Iniciado en el socket UNIX: %v, sistema operativo: %v",
			listener.Addr(),
			runtime.GOOS)

		defer listener.Close()

	default:
		log.Fatalf("Sistema operativo no soportado: %v", runtime.GOOS)
	}
}
