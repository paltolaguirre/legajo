package main

import (
	"log"
	"net/http"

	"github.com/xubiosueldos/conexionBD"
	"github.com/xubiosueldos/legajo/structLegajo"
)

func main() {
	tenant := "algo"
	db := conexionBD.ConnectBD(tenant)

	db.SingularTable(true)

	//para actualizar tablas...agrega columnas e indices, pero no elimina

	db.AutoMigrate(&structLegajo.Pais{}, &structLegajo.Provincia{}, &structLegajo.Localidad{}, &structLegajo.Zona{}, &structLegajo.Modalidadcontratacion{}, &structLegajo.Situacion{}, &structLegajo.Condicion{}, &structLegajo.Condicionsiniestrado{}, &structLegajo.Conveniocolectivo{}, &structLegajo.Centrodecosto{}, &structLegajo.Obrasocial{}, &structLegajo.Conyuge{}, &structLegajo.Hijo{}, &structLegajo.Legajo{})

	db.Model(&structLegajo.Hijo{}).AddForeignKey("legajoid", "legajo(id)", "NO ACTION", "NO ACTION")
	db.Model(&structLegajo.Conyuge{}).AddForeignKey("legajoid", "legajo(id)", "NO ACTION", "NO ACTION")

	//db.CreateTable(&structLegajo.Legajo{})

	router := newRouter()

	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)

}
