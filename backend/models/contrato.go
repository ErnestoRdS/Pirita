package models

import "gorm.io/gorm"

type Contrato struct {
	gorm.Model
	ConductorID uint `json:"-"`
	VehiculoID uint `json:"-"`
	FechaInicio string `json:"FechaInicio"`
	FechaFin string `json:"FechaFin"`
	Comisiones uint `json:"Comisiones"`
}