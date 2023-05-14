// El paquete models provee las estructuras de datos que se usan en el proyecto.
package models

import "gorm.io/gorm"

// Conductor es la estructura de datos que representa a un conductor en la base
// de datos.
type Conductor struct {
	// Incorporamos gorm.Model para que gorm sepa que debe usar el campo ID como
	// llave primaria, además de proveer los campos created_at y updated_at.
	gorm.Model

	// Nombre es el nombre del conductor.
	Nombre    string  `json:"nombre"`

	// Apellidos son los apellidos del conductor.
	Apellidos string  `json:"apellidos"`

	// CURP es el CURP del conductor.
	CURP      string  `json:"curp"`

	// ClaveINE es la clave de elector del conductor.
	ClaveINE  string  `json:"clave_ine"`

	// RFC es el RFC con homoclave del conductor.
	Salario   float64 `json:"salario"`

	// Estado es el estado en el que se encuentra el conductor. Puede ser
	// "activo", "inactivo" o "suspendido".
	Estado    string  `json:"estado"`

	// Pagos es la lista de pagos que ha recibido el conductor.
	Pagos     []Pago  `json:"pagos,omitempty" gorm:"foreignKey:ConductorID"`
}
