package structLegajo

import "github.com/xubiosueldos/conexionBD/structGormModel"

type Provincia struct {
	structGormModel.GormModel
	Nombre      string `json:"nombre"`
	Codigo      string `json:"codigo"`
	Descripcion string `json:"descripcion"`
	Activo      int    `json:"activo"`
	//Pais        Pais `gorm:"ForeignKey:Paisid;association_foreignkey:ID"`
	Paisid *uint `json:"paisid" sql:"type:int REFERENCES Pais(ID)"`
}
