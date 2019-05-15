package structLegajo

import "github.com/xubiosueldos/conexionBD/structGormModel"

type Localidad struct {
	structGormModel.GormModel
	Nombre      string `json:"nombre"`
	Codigo      string `json:"codigo"`
	Descripcion string `json:"descripcion"`
	Activo      int    `json:"activo"`
	//Provincia   Provincia `gorm:"ForeignKey:Provinciaid;association_foreignkey:ID"`
	Provinciaid *uint `json:"provinciaid" sql:"type:int REFERENCES Provincia(ID)"`
}
