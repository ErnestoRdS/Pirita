package models

import "gorm.io/gorm"

// La estructura Contrato representa un contrato entre un conductor y un vehículo.
type Contrato struct {
	// Incorporamos gorm.Model para que gorm sepa que debe usar el campo ID como
	// llave primaria, además de proveer los campos created_at y updated_at.
	gorm.Model

	// ConductorID es el ID del conductor al que pertenece el contrato. Es una
	// llave foránea.
	ConductorID uint   `json:"-"`

	// VehiculoID es el ID del vehículo al que pertenece el contrato. Es una
	// llave foránea.
	VehiculoID  uint   `json:"-"`

	// FechaInicio es la fecha en la que el contrato inicia.
	FechaInicio string `json:"FechaInicio"`

	// FechaFin es la fecha en la que el contrato termina.
	FechaFin    string `json:"FechaFin"`

	// Salario es el salario que recibe el conductor por el contrato, en este
	// caso la comisión es un porcentaje que el conductor recibe por cada viaje.
	Comisiones  uint   `json:"Comisiones"`
}
