package structLegajo

import (
	"time"

	"github.com/xubiosueldos/conexionBD/structGormModel"
)

type Legajo struct {
	structGormModel.GormModel
	Nombre                  string                 `json:"nombre"`
	Apellido                string                 `json:"apellido"`
	Codigo                  string                 `json:"codigo"`
	Descripcion             string                 `json:"descripcion"`
	Activo                  int                    `json:"activo"`
	Legajo                  string                 `json:"legajo"`
	Cuil                    string                 `json:"cuil"`
	Direccion               string                 `json:"direccion"`
	Localidad               *Localidad             `json:"localidad" gorm:"ForeignKey:Localidadid;association_foreignkey:ID;association_autoupdate:false"`
	Localidadid             *int                   `json:"localidadid" sql:"type:int REFERENCES Localidad(ID)"`
	Provincia               *Provincia             `json:"provincia" gorm:"ForeignKey:Provinciaid;association_foreignkey:ID;association_autoupdate:false"`
	Provinciaid             *int                   `json:"provinciaid" sql:"type:int REFERENCES Provincia(ID)"`
	Pais                    *Pais                  `json:"pais" gorm:"ForeignKey:Paisid;association_foreignkey:ID;association_autoupdate:false"`
	Paisid                  *int                   `json:"paisid" sql:"type:int REFERENCES Pais(ID)"`
	Zona                    *Zona                  `json:"zona" gorm:"ForeignKey:Zonaid;association_foreignkey:ID;association_autoupdate:false"`
	Zonaid                  *int                   `json:"zonaid" sql:"type:int REFERENCES Zona(ID)"`
	Telefono                string                 `json:"telefono"`
	Email                   string                 `json:"email"`
	Modalidadcontratacion   *Modalidadcontratacion `json:"modalidadcontratacion" gorm:"ForeignKey:Modalidadcontratacionid;association_foreignkey:ID;association_autoupdate:false"`
	Modalidadcontratacionid *int                   `json:"modalidadcontratacionid" sql:"type:int REFERENCES Modalidadcontratacion(ID)"`
	Categoria               string                 `json:"categoria"`
	Tarea                   string                 `json:"tarea"`
	Situacion               *Situacion             `json:"situacion" gorm:"ForeignKey:Situacionid;association_foreignkey:ID;association_autoupdate:false"`
	Situacionid             *int                   `json:"situacionid" sql:"type:int REFERENCES Situacion(ID)"`
	Condicion               *Condicion             `json:"condicion" gorm:"ForeignKey:Condicionid;association_foreignkey:ID;association_autoupdate:false"`
	Condicionid             *int                   `json:"condicionid" sql:"type:int REFERENCES Condicion(ID)"`
	Condicionsiniestrado    *Condicionsiniestrado  `json:"condicionsiniestrado" gorm:"ForeignKey:Condicionsiniestradoid;association_foreignkey:ID;association_autoupdate:false"`
	Condicionsiniestradoid  *int                   `json:"condicionsiniestradoid" sql:"type:int REFERENCES Condicionsiniestrado(ID)"`
	Obrasocial              *Obrasocial            `json:"obrasocial" gorm:"ForeignKey:Obrasocialid;association_foreignkey:ID;association_autoupdate:false"`
	Obrasocialid            *int                   `json:"obrasocialid" sql:"type:int REFERENCES Obrasocial(ID)"`
	Conveniocolectivo       *Conveniocolectivo     `json:"conveniocolectivo" gorm:"ForeignKey:Conveniocolectivoid;association_foreignkey:ID;association_autoupdate:false"`
	Conveniocolectivoid     *int                   `json:"conveniocolectivoid" sql:"type:int REFERENCES Conveniocolectivo(ID)"`
	Valorfijolrt            int                    `json:"valorfijolrt"`
	Conyuge                 []Conyuge              `json:"conyuge" gorm:"ForeignKey:Legajoid;association_foreignkey:ID"`
	Hijos                   []Hijo                 `json:"hijos" gorm:"ForeignKey:Legajoid;association_foreignkey:ID"`
	Remuneracion            int                    `json:"remuneracion"`
	HorasMensualesNormales  string                 `json:"horasmensualesnormales"`
	Fechaalta               time.Time              `json:"fechaalta"`
	Fechabaja               time.Time              `json:"fechabaja"`
	Centrodecosto           *Centrodecosto         `json:"centrodecosto" gorm:"ForeignKey:Centrodecostoid;association_foreignkey:ID;association_autoupdate:false"`
	Centrodecostoid         *int                   `json:"centrodecostoid" sql:"type:int REFERENCES Centrodecosto(ID)"`
	Cbu                     string                 `json:"cbu"`
}
