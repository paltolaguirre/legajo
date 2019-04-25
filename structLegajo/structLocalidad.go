package structLegajo

import (
	"github.com/jinzhu/gorm"
)

type Localidad struct {
	gorm.Model
	Nombre      string `json:"nombre"`
	Codigo      string `json:"codigo"`
	Descripcion string `json:"descripcion"`
	Activo      int    `json:"activo"`
	//Provincia   Provincia `gorm:"ForeignKey:Provinciaid;association_foreignkey:ID"`
	Provinciaid uint `json:"provinciaid" sql:"type:int REFERENCES Provincia(ID)"`
}
