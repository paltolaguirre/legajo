package main

import (
  	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"

)

func connectBD()(*gorm.DB){

	db, err := gorm.Open("postgres", "host=192.168.30.111 port=5432 user=postgres dbname=DES_MULTITENANT_AR_1 password=Post66MM/")

	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()
	return db

}
