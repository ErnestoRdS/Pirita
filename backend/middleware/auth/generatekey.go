// El paquete auth contiene los middlewares que se usan en la aplicación
// de Fiber para manejar la autenticación de rutas con API keys.
package auth

import (
	"github.com/google/uuid"
)

// GenerateKey genera una API key aleatoria con un UUID y la devuelve como string.
func GenerateKey() string {
	return uuid.NewString()
}
