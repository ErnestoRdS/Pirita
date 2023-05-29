package routes

import (
	"os"
	"time"

	"github.com/UpVent/Pirita/v2/models"
	"github.com/gofiber/fiber/v2"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"

	"gorm.io/gorm"
)

// LoginRouter es router para la autenticación de usuarios en Pirita.
func LoginRouter(app *fiber.App, db *gorm.DB) {
	app.Post("/login", func(c *fiber.Ctx) error {

		type LoginInput struct {
			Usuario  string `json:"usuario"`
			Password string `json:"password"`
		}

		input := new(LoginInput)

		if err := c.BodyParser(input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Error al decodificar el cuerpo de la petición",
			})
		}

		var admin models.Administrador
		var conductor models.Conductor

		// Buscar al usuario en la base de datos.
		if err := db.Where("usuario = ?", input.Usuario).First(&admin).Error; err != nil {
			if err := db.Where("usuario = ?", input.Usuario).First(&conductor).Error; err != nil {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"message": "Usuario no encontrado.",
				})
			}
		}

		// Comprobar la contraseña.
		if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(input.Password)); err != nil {
			if err := bcrypt.CompareHashAndPassword([]byte(conductor.Password), []byte(input.Password)); err != nil {
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

		secret := []byte(os.Getenv("JWT_SECRET"))

		// Firmar el token.
		tokenString, err := token.SignedString(secret)

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
