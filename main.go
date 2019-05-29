package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/autenticacion/publico"
	"github.com/xubiosueldos/conexionBD/apiclientconexionbd"
	"github.com/xubiosueldos/framework/configuracion"
	"github.com/xubiosueldos/legajo/structLegajo"
)

func main() {
	configuracion := configuracion.GetInstance()
	var tokenAutenticacion publico.Security
	tokenAutenticacion.Tenant = "public"

	apiclientconexionbd.ObtenerDB(&tokenAutenticacion, nombreMicroservicio, obtenerVersionLegajo(), AutomigrateTablasPublicas)

	router := newRouter()

	server := http.ListenAndServe(":"+configuracion.Puertomicroserviciolegajo, router)

	log.Fatal(server)

}

func AutomigrateTablasPublicas(db *gorm.DB) {

	//para actualizar tablas...agrega columnas e indices, pero no elimina
	db.AutoMigrate(&structLegajo.Pais{}, &structLegajo.Provincia{}, &structLegajo.Localidad{}, &structLegajo.Zona{}, &structLegajo.Modalidadcontratacion{}, &structLegajo.Situacion{}, &structLegajo.Condicion{}, &structLegajo.Condicionsiniestrado{}, &structLegajo.Conveniocolectivo{}, &structLegajo.Centrodecosto{}, &structLegajo.Obrasocial{})

}
