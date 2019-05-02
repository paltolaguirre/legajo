package main

import (
	"log"
	"net/http"

	"github.com/xubiosueldos/conexionBD"
	"github.com/xubiosueldos/legajo/structLegajo"
)

func main() {
	var tenant string = "public"
	db := conexionBD.ConnectBD(tenant)

	//para actualizar tablas...agrega columnas e indices, pero no elimina
	db.AutoMigrate(&structLegajo.Pais{}, &structLegajo.Provincia{}, &structLegajo.Localidad{}, &structLegajo.Zona{}, &structLegajo.Modalidadcontratacion{}, &structLegajo.Situacion{}, &structLegajo.Condicion{}, &structLegajo.Condicionsiniestrado{}, &structLegajo.Conveniocolectivo{}, &structLegajo.Centrodecosto{}, &structLegajo.Obrasocial{})

	router := newRouter()

	server := http.ListenAndServe(":8083", router)

	log.Fatal(server)

}
