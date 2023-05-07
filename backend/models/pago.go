package models

import "gorm.io/gorm"

type Pago struct {
	gorm.Model
	ConductorID uint `json:"-"`
	Fecha string `json:"fecha"`
	Cantidad float64 `json:"cantidad"`
	Notas string `json:"nota"`
}
