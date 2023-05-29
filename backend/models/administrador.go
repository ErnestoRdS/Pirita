package models

import "gorm.io/gorm"

// Administrador es el modelo de datos para los administradores de la aplicación.
// Los administradores son los encargados de gestionar los conductores, así como
// de gestionar los pagos que se les hacen a los conductores
type Administrador struct {
	// Incorporamos gorm.Model para que gorm sepa que debe usar el campo ID como
	// llave primaria, además de proveer los campos created_at y updated_at.
	gorm.Model

	// Nombre es el nombre del Administrador.
	Nombre string `json:"nombre"`

	// Apellidos son los apellidos del Administrador.
	Apellidos string `json:"apellidos"`

	// Para el inicio de sesión.

	// Usuario es el nombre de usuario del Administrador.
	Usuario string `json:"usuario"`

	// Correo es el correo electrónico del Adminisrador.
	Correo string `json:"correo"`

	// Password es la contraseña del Administrador.
	Password string `json:"password"`
}
