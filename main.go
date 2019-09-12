package main

import (
	"log"
	"net/http"

	"github.com/xubiosueldos/autenticacion/apiclientautenticacion"
	"github.com/xubiosueldos/conexionBD/Autenticacion/structAutenticacion"
	"github.com/xubiosueldos/conexionBD/apiclientconexionbd"
	"github.com/xubiosueldos/framework/configuracion"
)

func main() {
	configuracion := configuracion.GetInstance()
	var tokenAutenticacion structAutenticacion.Security
	tokenAutenticacion.Tenant = "public"

	tenant := apiclientautenticacion.ObtenerTenant(&tokenAutenticacion)
	apiclientconexionbd.ObtenerDB(tenant, nombreMicroservicio, obtenerVersionLegajo())

	router := newRouter()

	server := http.ListenAndServe(":"+configuracion.Puertomicroserviciolegajo, router)

	log.Fatal(server)

}
