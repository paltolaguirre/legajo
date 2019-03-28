package structLegajo

import (
	"github.com/jinzhu/gorm"
)

type Conyuge struct {
	gorm.Model
	Nombre       string
	Apellido     string
	Codigo       string
	Descripcion  string
	Activo       int
	Cuil         string
	Obrasocial   Obrasocial `gorm:"ForeignKey:Obrasocialid;association_foreignkey:ID"`
	Obrasocialid uint       `sql:"type:int REFERENCES Obrasocial(ID)"`
	Legajoid     uint
}
