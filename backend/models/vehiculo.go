package models

import "gorm.io/gorm"

// Vehiculo es la estructura de datos que representa un vehiculo en la base de datos.
type Vehiculo struct {
	// Incorporamos gorm.Model para que gorm sepa que debe usar el campo ID como
	// llave primaria, además de proveer los campos created_at y updated_at.
	gorm.Model

	// Fabricante es el fabricante del vehiculo.
	Fabricante string `json:"Fabricante"`

	// Marca es la marca del vehiculo.
	Marca string `json:"Marca"`

	// Modelo es el modelo del vehiculo.
	Modelo uint `json:"Modelo"`

	// Placas son las placas del vehiculo.
	Placas string `json:"Placas"`

	// Color es el color del vehiculo.
	Color string `json:"Color"`

	// VigenciaTec es la vigencia de la tarjeta de circulación del vehiculo.
	VigenciaTec string `json:"VigenciaTec"`

	// VigenciaSeg es la vigencia del seguro del vehiculo.
	Seguro bool `json:"Seguro"`

	// Estado es el estado en el que se encuentra el vehiculo. Puede ser
	// "activo", "inactivo" o "suspendido".
	Estado string `json:"Estado"`
}
