package structLegajo

import (
	"github.com/jinzhu/gorm"
)

type Localidad struct {
	gorm.Model
	Nombre      string
	Codigo      string
	Descripcion string
	Activo      int
	//Provincia   Provincia `gorm:"ForeignKey:Provinciaid;association_foreignkey:ID"`
	Provinciaid uint `sql:"type:int REFERENCES Provincia(ID)"`
}
