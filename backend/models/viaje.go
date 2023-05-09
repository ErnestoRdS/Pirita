package models

import "gorm.io/gorm"

type Viaje struct {
	gorm.Model
	ConductorID uint    `json:"-"`
	VehiculoID  uint    `json:"-"`
	Fecha       string  `json:"Fecha"`
	Origen      string  `json:"Origen"`
	Destino     string  `json:"Destino"`
	Monto       float64 `json:"Monto"`
}
