package models

import "gorm.io/gorm"

// Pago es la estructura de datos que representa un pago en la base de datos.
type Pago struct {
	// Incorporamos gorm.Model para que gorm sepa que debe usar el campo ID como
	// llave primaria, además de proveer los campos created_at y updated_at.
	gorm.Model

	// ConductorID es el ID del conductor al que pertenece el pago. Es una llave
	// foránea.
	ConductorID uint `json:"-"`

	// Fecha es la fecha en la que se realizó el pago.
	Fecha string `json:"fecha"`

	// Cantidad es la cantidad que se pagó al conductor.
	Cantidad float64 `json:"cantidad"`

	// Notas es una nota que se puede agregar al pago.
	Notas string `json:"nota"`
}
