// Pirita Backend es un binario hecho en Go que sirve como backend para la
// plataforma de Pirita. Este binario se encarga de servir las rutas protegidas
// de la API, así como las rutas públicas de la misma. Maneja la autenticación
// y autorización de los usuarios, así como la creación de los mismos.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"

	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	flogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	"github.com/google/uuid"

	jwtware "github.com/gofiber/contrib/jwt"

	"github.com/UpVent/Pirita/v2/middleware"
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

	// Verificar si el archivo de variables de entorno existe.
	_, err := os.Stat(".env")

	// Si el archivo no existe, crearlo y generar un uuid para la clave JWT.
	if os.IsNotExist(err) {
		secret := uuid.New().String()
		err = ioutil.WriteFile(".env", []byte(fmt.Sprintf("JWT_SECRET=%s", secret)), 0644)

		if err != nil {
			log.Fatalf("Error al crear el archivo de variables de entorno: %v", err)
		}
	}

	err = godotenv.Load()

	if err != nil {
		log.Fatalf("Error al cargar las variables de entorno: %v", err)
	}

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
		&models.Administrador{},
		&models.Conductor{},
		&models.Contrato{},
		&models.Pago{},
		&models.Vehiculo{},
		&models.Viaje{},
	)

	if err != nil {
		log.Fatalf("Error al migrar los modelos a la base de datos: %v", err)
	}

	var admin models.Administrador
	if err := db.First(&admin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			DefaultAdminPassword := "PiritaAdmin"
			DefaultUserPassword := "conductordemo"

			AdminHashedPassword, err := bcrypt.GenerateFromPassword([]byte(DefaultAdminPassword), bcrypt.DefaultCost)

			if err != nil {
				log.Fatalf("Error al generar la contraseña aleatoria para el administrador: %v", err)
			}

			UserHashedPassword, err := bcrypt.GenerateFromPassword([]byte(DefaultUserPassword), bcrypt.DefaultCost)

			if err != nil {
				log.Fatalf("Error al generar la contraseña aleatoria para el usuario común: %v", err)
			}

			// Crear registros de prueba utilizando GORM
			admin := models.Administrador{
				Nombre:    "Administrador",
				Apellidos: "Pirita",
				Usuario:   "admin",
				Correo:    "admin@pirita.com",
				Password:  string(AdminHashedPassword),
			}

			db.Create(&admin)

			conductor := models.Conductor{
				Nombre:     "Conductor",
				Apellidos:  "Pirita",
				CURP:       "PIRP-00000000000000",
				ClaveINE:   "PIRP-00000000000000",
				RFC:        "PIRP-00000000000000",
				Salario:    50.00,
				Comisiones: 10,
				Estado:     "Activo",
				Usuario:    "conductor",
				Correo:     "conductor@pirita.com",
				Password:   string(UserHashedPassword),
			}

			db.Create(&conductor)

			vehiculo := models.Vehiculo{
				Fabricante:  "Bentley",
				Marca:       "Continental GT",
				Modelo:      2015,
				Placas:      "FUK-04-EVA",
				Color:       "manzana",
				VigenciaTec: "28-05-2023",
				Seguro:      true,
				Estado:      "Activo",
			}
			db.Create(&vehiculo)

			contrato := models.Contrato{
				ConductorID: conductor.ID,
				VehiculoID:  vehiculo.ID,
				FechaInicio: "28-05-2023",
				FechaFin:    "28-06-2023",
				Comisiones:  10,
			}
			db.Create(&contrato)

			viaje := models.Viaje{
				ConductorID: conductor.ID,
				VehiculoID:  vehiculo.ID,
				Fecha:       "28-05-2023",
				Monto:       50.0,
			}
			db.Create(&viaje)

			pago := models.Pago{
				ConductorID: conductor.ID,
				Fecha:       "28-05-2023",
				Cantidad:    15.0,
				Notas:       "Sin comentarios",
			}
			db.Create(&pago)

			// Imprimir los detalles de la cuenta de administrador en la terminal.
			fmt.Printf("Cuenta de administrador creada con éxito.\nUsuario: %s\nContraseña: %s\n", admin.Usuario, DefaultAdminPassword)
			fmt.Printf("Cuenta de usuario común creada con éxito.\nUsuario: %s\nContraseña: %s\n", conductor.Usuario, DefaultUserPassword)
		} else {
			log.Fatalf("Error al buscar la cuenta de administrador: %v", err)
		}
	}

	// Crear la aplicación de Fiber.
	app := fiber.New()

	// Middelewares.
	// Usar el middleware de logger y el de limitador de peticiones.
	app.Use(flogger.New())
	app.Use(limiter.New())
	app.Use(cache.New())

	// Usar el middleware contrib de JWT.
	jwtMiddleware := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
	})

	// Montar las rutas.
	routes.ConductorRouter(app, db, jwtMiddleware)

	routes.ContratoRouter(app, db, jwtMiddleware)

	routes.PagosRouter(app, db, jwtMiddleware)

	routes.VehiculoRouter(app, db, jwtMiddleware)

	routes.ViajeRouter(app, db, jwtMiddleware)

	routes.LoginRouter(app, db)

	// Ruta de monitoreo.
	app.Use("/monitor", jwtMiddleware, middleware.AdminMiddleware)
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
	log.Fatal(app.Listen(":" + port))
}
