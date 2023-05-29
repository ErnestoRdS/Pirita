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


	randomPassword := "test1234"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(randomPassword), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("Error al generar la contraseña aleatoria: %v", err)
	}

	log.Printf("Guarde bien esta contraseña, es la única vez que la verá: %s", randomPassword)

	admin := models.Administrador{
		Nombre:   "Administrador",
		Apellidos: "Pirita",
		Usuario:  "admin",
		Correo:   "admin@pirita.com",
		Password: string(hashedPassword),
	}

	db.Create(&admin)


	// Sentencias para llenar la base de datos con datos de prueba
	// Un registro de prueba para cada tabla
	states := `
		INSERT INTO conductors(nombre, apellidos, curp, clave_ine, salario, estado) VALUES('Yo', ':3xd', ':3xY666666HCSFUK02', 'claveineYoxd', 150.0, 'Pasivo');
		INSERT INTO vehiculos(fabricante, marca, modelo, placas, color, vigencia_tec, seguro, estado) VALUES('Bentley', 'Continental GT', 2015, 'FUK-04-EVA', 'manzana', 'Simónxd', 'Tambiénxd', 'Sin defensa trasera');                
		INSERT INTO contratos(conductor_id, vehiculo_id, fecha_inicio, fecha_fin, comisiones) VALUES(1, 1, '28-05-2023', '28-06-2023', 10);
		INSERT INTO viajes(conductor_id, vehiculo_id, fecha, monto) VALUES(1, 1, '28-05-2023', 50.0);
		INSERT INTO pagos(conductor_id, fecha, cantidad, notas) VALUES(1, '28-05-2023', 15.0, 'Sin comentarios');
	`

	// Ejecutar sentencias con cadenas de SQL en crudo
	db.Exec(states)

	// Crear la aplicación de Fiber.
	app := fiber.New()

	// Middelewares.
	// Usar el middleware de logger y el de limitador de peticiones.
	app.Use(flogger.New())
	app.Use(limiter.New())
	app.Use(cache.New())

	jwtMiddleware := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
	})

	// Montar las rutas.
	routes.ConductorRouter(app, db, jwtMiddleware)

	app.Use("/api/contratos", jwtMiddleware)
	routes.ContratoRouter(app, db)

	app.Use("/api/pagos", jwtMiddleware)
	routes.PagosRouter(app, db)

	app.Use("/api/vehiculos", jwtMiddleware)
	routes.VehiculoRouter(app, db)

	app.Use("/api/viajes", jwtMiddleware)
	routes.ViajeRouter(app, db)

	routes.LoginRouter(app, db)

	// Ruta de monitoreo.
	app.Use("/monitor", jwtMiddleware)
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
