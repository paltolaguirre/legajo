package structLegajo

import (
	"time"

	"github.com/xubiosueldos/conexionBD/structGormModel"
)

type Legajo struct {
	structGormModel.GormModel
	Nombre                  *string                `json:"nombre" gorm:"not null"`
	Apellido                *string                `json:"apellido" gorm:"not null"`
	Codigo                  string                 `json:"codigo"`
	Descripcion             string                 `json:"descripcion"`
	Activo                  int                    `json:"activo"`
	Legajo                  string                 `json:"legajo"`
	Cuil                    *string                `json:"cuil" gorm:"not null"`
	Direccion               string                 `json:"direccion"`
	Localidad               *Localidad             `json:"localidad" gorm:"ForeignKey:Localidadid;association_foreignkey:ID;association_autoupdate:false;not null"`
	Localidadid             *int                   `json:"localidadid" sql:"type:int REFERENCES Localidad(ID)" gorm:"not null"`
	Provincia               *Provincia             `json:"provincia" gorm:"ForeignKey:Provinciaid;association_foreignkey:ID;association_autoupdate:false;not null"`
	Provinciaid             *int                   `json:"provinciaid" sql:"type:int REFERENCES Provincia(ID)" gorm:"not null"`
	Pais                    *Pais                  `json:"pais" gorm:"ForeignKey:Paisid;association_foreignkey:ID;association_autoupdate:false;not null"`
	Paisid                  *int                   `json:"paisid" sql:"type:int REFERENCES Pais(ID)" gorm:"not null"`
	Zona                    *Zona                  `json:"zona" gorm:"ForeignKey:Zonaid;association_foreignkey:ID;association_autoupdate:false"`
	Zonaid                  *int                   `json:"zonaid" sql:"type:int REFERENCES Zona(ID)"`
	Telefono                string                 `json:"telefono"`
	Email                   string                 `json:"email"`
	Modalidadcontratacion   *Modalidadcontratacion `json:"modalidadcontratacion" gorm:"ForeignKey:Modalidadcontratacionid;association_foreignkey:ID;association_autoupdate:false;not null"`
	Modalidadcontratacionid *int                   `json:"modalidadcontratacionid" sql:"type:int REFERENCES Modalidadcontratacion(ID)" gorm:"not null"`
	Categoria               *string                `json:"categoria" gorm:"not null"`
	Tarea                   *string                `json:"tarea" gorm:"not null"`
	Situacion               *Situacion             `json:"situacion" gorm:"ForeignKey:Situacionid;association_foreignkey:ID;association_autoupdate:false;not null"`
	Situacionid             *int                   `json:"situacionid" sql:"type:int REFERENCES Situacion(ID)" gorm:"not null"`
	Condicion               *Condicion             `json:"condicion" gorm:"ForeignKey:Condicionid;association_foreignkey:ID;association_autoupdate:false;not null"`
	Condicionid             *int                   `json:"condicionid" sql:"type:int REFERENCES Condicion(ID)" gorm:"not null"`
	Condicionsiniestrado    *Condicionsiniestrado  `json:"condicionsiniestrado" gorm:"ForeignKey:Condicionsiniestradoid;association_foreignkey:ID;association_autoupdate:false;not null"`
	Condicionsiniestradoid  *int                   `json:"condicionsiniestradoid" sql:"type:int REFERENCES Condicionsiniestrado(ID)" gorm:"not null"`
	Obrasocial              *Obrasocial            `json:"obrasocial" gorm:"ForeignKey:Obrasocialid;association_foreignkey:ID;association_autoupdate:false;not null"`
	Obrasocialid            *int                   `json:"obrasocialid" sql:"type:int REFERENCES Obrasocial(ID)" gorm:"not null"`
	Conveniocolectivo       *Conveniocolectivo     `json:"conveniocolectivo" gorm:"ForeignKey:Conveniocolectivoid;association_foreignkey:ID;association_autoupdate:false;not null"`
	Conveniocolectivoid     *int                   `json:"conveniocolectivoid" sql:"type:int REFERENCES Conveniocolectivo(ID)" gorm:"not null"`
	Valorfijolrt            int                    `json:"valorfijolrt"`
	Conyuge                 []Conyuge              `json:"conyuge" gorm:"ForeignKey:Legajoid;association_foreignkey:ID"`
	Hijos                   []Hijo                 `json:"hijos" gorm:"ForeignKey:Legajoid;association_foreignkey:ID"`
	Remuneracion            int                    `json:"remuneracion" gorm:"not null"`
	Horasmensualesnormales  string                 `json:"horasmensualesnormales" gorm:"not null"`
	Fechaalta               *time.Time             `json:"fechaalta" gorm:"not null"`
	Fechabaja               *time.Time             `json:"fechabaja"`
	Centrodecosto           *Centrodecosto         `json:"centrodecosto" gorm:"ForeignKey:Centrodecostoid;association_foreignkey:ID;association_autoupdate:false"`
	Centrodecostoid         *int                   `json:"centrodecostoid" sql:"type:int REFERENCES Centrodecosto(ID)"`
	Cbu                     string                 `json:"cbu"`
}
