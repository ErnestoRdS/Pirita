package models

import "gorm.io/gorm"

type Vehiculo struct {
	gorm.Model
	Fabricante  string `json:"Fabricante"`
	Marca       string `json:"Marca"`
	Modelo      uint   `json:"Modelo"`
	Placas      string `json:"Placas"`
	Color       string `json:"Color"`
	VigenciaTec string `json:"VigenciaTec"`
	Seguro      bool   `json:"Seguro"`
	Estado      string `json:"Estado"`
}
