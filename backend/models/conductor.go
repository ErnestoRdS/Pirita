// El paquete models provee las estructuras de datos que se usan en el proyecto.
package models

import "gorm.io/gorm"

type Conductor struct {
	gorm.Model
	Nombre string `json:"nombre"`
	Apellidos string `json:"apellidos"`
	CURP string `json:"curp"`
	ClaveINE string `json:"clave_ine"`
	Salario float64	`json:"salario"`
	Estado    string  `json:"estado"`
	Pagos []Pago `json:"pagos,omitempty" gorm:"foreignKey:ConductorID"`
}
