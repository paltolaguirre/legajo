package structLegajo

import (
	"github.com/jinzhu/gorm"
)

type Provincia struct {
	gorm.Model
	Nombre      string
	Codigo      string
	Descripcion string
	Activo      int
	Pais        Pais `gorm:"ForeignKey:Paisid;association_foreignkey:ID"`
	Paisid      uint `sql:"type:int REFERENCES Pais(ID)"`
}
