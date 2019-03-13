package main

import ("net/http"
 		"log"
 		"github.com/xubiosueldos/conexionBD"
 		"github.com/xubiosueldos/legajo/structLegajo"
)

func main() {
	tenant := "algo"
	db := conexionBD.ConnectBD(tenant)
	db.CreateTable(&structLegajo.Legajo{})

	router := newRouter()
	
	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)

}
