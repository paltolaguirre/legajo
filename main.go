package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/xubiosueldos/framework/configuracion"
)

func main() {
	configuracion := configuracion.GetInstance()
	router := newRouter()

	fmt.Println("Microservicio de Legajo en el puerto: " + configuracion.Puertomicroserviciolegajo)
	server := http.ListenAndServe(":"+configuracion.Puertomicroserviciolegajo, router)

	log.Fatal(server)

}
