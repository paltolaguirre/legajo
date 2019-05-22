package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/autenticacion/publico"
	"github.com/xubiosueldos/conexionBD/apiclientconexionbd"
	"github.com/xubiosueldos/legajo/structLegajo"
)

func main() {
	var tokenAutenticacion publico.Security
	tokenAutenticacion.Tenant = "public"

	apiclientconexionbd.ObtenerDB(&tokenAutenticacion, "legajo", 1, AutomigrateTablasPublicas)

	router := newRouter()

	server := http.ListenAndServe(":8083", router)

	log.Fatal(server)

}

func AutomigrateTablasPublicas(db *gorm.DB) {

	//para actualizar tablas...agrega columnas e indices, pero no elimina
	db.AutoMigrate(&structLegajo.Pais{}, &structLegajo.Provincia{}, &structLegajo.Localidad{}, &structLegajo.Zona{}, &structLegajo.Modalidadcontratacion{}, &structLegajo.Situacion{}, &structLegajo.Condicion{}, &structLegajo.Condicionsiniestrado{}, &structLegajo.Conveniocolectivo{}, &structLegajo.Centrodecosto{}, &structLegajo.Obrasocial{})

}
