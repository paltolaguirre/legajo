package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/autenticacion/apiclientautenticacion"
	"github.com/xubiosueldos/autenticacion/publico"
	"github.com/xubiosueldos/conexionBD/apiclientconexionbd"
	"github.com/xubiosueldos/framework/configuracion"
	"github.com/xubiosueldos/legajo/structLegajo"
)

func main() {
	configuracion := configuracion.GetInstance()
	var tokenAutenticacion publico.Security
	tokenAutenticacion.Tenant = "public"

	tenant := apiclientautenticacion.ObtenerTenant(&tokenAutenticacion)
	apiclientconexionbd.ObtenerDB(tenant, nombreMicroservicio, obtenerVersionLegajo(), AutomigrateTablasPublicas)

	router := newRouter()

	server := http.ListenAndServe(":"+configuracion.Puertomicroserviciolegajo, router)

	log.Fatal(server)

}

func AutomigrateTablasPublicas(db *gorm.DB) {

	//para actualizar tablas...agrega columnas e indices, pero no elimina
	db.AutoMigrate(&structLegajo.Pais{}, &structLegajo.Provincia{}, &structLegajo.Localidad{}, &structLegajo.Modalidadcontratacion{}, &structLegajo.Situacion{}, &structLegajo.Condicion{}, &structLegajo.Condicionsiniestrado{}, &structLegajo.Obrasocial{})

}
