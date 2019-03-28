package structLegajo

import (
	"github.com/jinzhu/gorm"
)

type Zona struct {
	gorm.Model
	Nombre      string
	Codigo      string
	Descripcion string
	Activo      int
}
