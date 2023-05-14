package models

import "gorm.io/gorm"

// Viaje es la estructura de datos que representa un viaje en la base de datos.
type Viaje struct {
	// Incorporamos gorm.Model para que gorm sepa que debe usar el campo ID como
	// llave primaria, además de proveer los campos created_at y updated_at.
	gorm.Model

	// ConductorID es el ID del conductor al que pertenece el viaje. Es una llave
	// foránea.
	ConductorID uint    `json:"-"`

	// VehiculoID es el ID del vehículo al que pertenece el viaje. Es una llave
	// foránea.
	VehiculoID  uint    `json:"-"`

	// Fecha es la fecha en la que se realizó el viaje.
	Fecha       string  `json:"Fecha"`

	// Origen es el origen del viaje.
	Monto       float64 `json:"Monto"`
}
