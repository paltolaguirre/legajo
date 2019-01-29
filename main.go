package main

import "net/http"
import "log"

func main() {

	db := connectBD()
	db.CreateTable(&Legajo{})

	router := newRouter()
	
	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)

}
