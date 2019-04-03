package structLegajo

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Legajo struct {
	gorm.Model
	Nombre                  string
	Apellido                string
	Codigo                  string
	Descripcion             string
	Activo                  int
	Legajo                  string
	Cuil                    string
	Direccion               string
	Localidad               Localidad `gorm:"ForeignKey:Localidadid;association_foreignkey:ID;association_autoupdate:false"`
	Localidadid             uint      `sql:"type:int REFERENCES Localidad(ID)"`
	Provincia               Provincia `gorm:"ForeignKey:Provinciaid;association_foreignkey:ID;association_autoupdate:false"`
	Provinciaid             uint      `sql:"type:int REFERENCES Provincia(ID)"`
	Pais                    Pais      `gorm:"ForeignKey:Paisid;association_foreignkey:ID;association_autoupdate:false"`
	Paisid                  uint      `sql:"type:int REFERENCES Pais(ID)"`
	Zona                    Zona      `gorm:"ForeignKey:Zonaid;association_foreignkey:ID;association_autoupdate:false"`
	Zonaid                  uint      `sql:"type:int REFERENCES Zona(ID)"`
	Telefono                string
	Email                   string
	Modalidadcontratacion   Modalidadcontratacion `gorm:"ForeignKey:Modalidadcontratacionid;association_foreignkey:ID;association_autoupdate:false"`
	Modalidadcontratacionid uint                  `sql:"type:int REFERENCES Modalidadcontratacion(ID)"`
	Categoria               string
	Tarea                   string
	Situacion               Situacion            `gorm:"ForeignKey:Situacionid;association_foreignkey:ID;association_autoupdate:false"`
	Situacionid             uint                 `sql:"type:int REFERENCES Situacion(ID)"`
	Condicion               Condicion            `gorm:"ForeignKey:Condicionid;association_foreignkey:ID;association_autoupdate:false"`
	Condicionid             uint                 `sql:"type:int REFERENCES Condicion(ID)"`
	Condicionsiniestrado    Condicionsiniestrado `gorm:"ForeignKey:Condicionsiniestradoid;association_foreignkey:ID;association_autoupdate:false"`
	Condicionsiniestradoid  uint                 `sql:"type:int REFERENCES Condicionsiniestrado(ID)"`
	Obrasocial              Obrasocial           `gorm:"ForeignKey:Obrasocialid;association_foreignkey:ID;association_autoupdate:false"`
	Obrasocialid            uint                 `sql:"type:int REFERENCES Obrasocial(ID)"`
	Conveniocolectivo       Conveniocolectivo    `gorm:"ForeignKey:Conveniocolectivoid;association_foreignkey:ID;association_autoupdate:false"`
	Conveniocolectivoid     uint                 `sql:"type:int REFERENCES Conveniocolectivo(ID)"`
	Valorfijolrt            int
	Conyuge                 Conyuge `gorm:"ForeignKey:Legajoid;association_foreignkey:ID"`
	Hijos                   []Hijo  `gorm:"ForeignKey:Legajoid;association_foreignkey:ID"`
	Remuneracion            int
	HorasMensualesNormales  string
	Fechaalta               time.Time
	Fechabaja               time.Time
	Centrodecosto           Centrodecosto `gorm:"ForeignKey:Centrodecostoid;association_foreignkey:ID;association_autoupdate:false"`
	Centrodecostoid         uint          `sql:"type:int REFERENCES Centrodecosto(ID)"`
	Cbu                     string
}
