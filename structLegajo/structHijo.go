package structLegajo

import "github.com/xubiosueldos/conexionBD/structGormModel"

type Hijo struct {
	structGormModel.GormModel
	Nombre       string      `json:"nombre"`
	Apellido     string      `json:"apellido"`
	Codigo       string      `json:"codigo"`
	Descripcion  string      `json:"descripcion"`
	Activo       int         `json:"activo"`
	Cuil         string      `json:"cuil"`
	Obrasocial   *Obrasocial `json:"obrasocial" gorm:"ForeignKey:Obrasocialid;association_foreignkey:ID"`
	Obrasocialid *int        `json:"obrasocialid" sql:"type:int REFERENCES Obrasocial(ID)"`
	Legajoid     *int        `json:"legajoid"` //`sql:"type:int REFERENCES Legajo(ID)"`
}
