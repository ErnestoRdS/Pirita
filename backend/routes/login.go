package routes

import (
	"os"
	"time"

	"github.com/UpVent/Pirita/v2/models"
	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt/v5"

	"gorm.io/gorm"
)

// LoginRouter es router para la autenticación de usuarios en Pirita.
func LoginRouter(app *fiber.App, db *gorm.DB) {
	app.Get("/login", func(c *fiber.Ctx) error {
		user := c.FormValue("user", "")
		pass := c.FormValue("pass", "")

		var admin models.Administrador
		var conductor models.Conductor

		// Buscar al usuario en la base de datos.
		if err := db.Where("usuario = ?", user).First(&admin).Error; err != nil {
			if err := db.Where("usuario = ?", user).First(&conductor).Error; err != nil {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"message": "Usuario no encontrado.",
				})
			}
		}

		// Comprobar la contraseña.
		if admin.Password != pass {
			if conductor.Password != pass {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Contraseña incorrecta.",
				})
			}
		}

		// Crear el token.
		token := jwt.New(jwt.SigningMethodHS256)

		// Añadir los claimms dependiendo del tipo de usuario.
		if admin.ID != 0 {
			claims := token.Claims.(jwt.MapClaims)
			claims["id"] = admin.ID
			claims["user"] = admin.Usuario
			claims["type"] = "admin"
			claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		} else {
			claims := token.Claims.(jwt.MapClaims)
			claims["id"] = conductor.ID
			claims["user"] = conductor.Usuario
			claims["type"] = "conductor"
			claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		}


		// Firmar el token.
		tokenString, err := token.SignedString(os.Getenv("JWT_SECRET"))

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error al firmar el token.",
			})
		}

		// Devolver el token.
		return c.JSON(fiber.Map{
			"token": tokenString,
		})
	})
}
