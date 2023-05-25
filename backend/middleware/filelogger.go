// El paquete middleware contiene los middlewares que se usan en la aplicación
// de Fiber.
package middleware

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// FileLogger es un middleware que se encarga de escribir los logs de la
// aplicación en un archivo de texto, este archivo se encuentra en la carpeta
// temporal del sistema, y se llama pirita.log.
func FileLogger() fiber.Handler {
	// Crear el archivo de logs en la carpeta temporal del sistema.
	tmpDir := os.TempDir()

	logfile, err := os.OpenFile(tmpDir+"/pirita.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Error al crear el archivo de logs: %v", err)
		log.Fatal("No se pudo crear el archivo de logs, el sistema continuará trabajando pero no se escribirán logs.")
	}

	defer logfile.Close()

	return logger.New(logger.Config{
		Output: logfile,
	})
}
