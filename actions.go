package main

import (
	"fmt"
    "net/http"
    "log"
    "github.com/gorilla/mux"
    "encoding/json"
  	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"

)

var db *gorm.DB
var err error

type Legajo struct {
	gorm.Model
	Nombre string
	Codigo string
	Descripcion string
	Activo int
}

type Legajos []Legajo


/*func connectBD()(*gorm.DB){

	db, err = gorm.Open("postgres", "host=192.168.30.111 port=5432 user=postgres dbname=DES_MULTITENANT_AR_1 password=Post66MM/")

	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()
	return db

}*/


func responseLegajo(w http.ResponseWriter, status int, results Legajo){

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(results)
}


func responseLegajos(w http.ResponseWriter, status int, results []Legajo){

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(results)

}


func Index(w http.ResponseWriter, r *http.Request){
  
	fmt.Fprintf(w, "Probando API Rest con Gorm")

}


func LegajoList(w http.ResponseWriter, r *http.Request){

	db := connectBD()

	var legajos []Legajo

	//Lista todos los legajos 
	fmt.Println("Los legajos de la BD son: ")
	

	db.Find(&legajos)

	fmt.Println(legajos)
	responseLegajos(w, 202, legajos)

}

func LegajoShow(w http.ResponseWriter, r *http.Request){

	params := mux.Vars(r)
	legajo_id := params["id"]

	fmt.Println(legajo_id)

	var legajo Legajo	//Con &var --> lo que devuelve el metodo se le asigna a la var

	//db.First(&legajo, "id = ?", legajo_id)

	asd := connectBD()
	asd.First(&legajo, "id = ?", legajo_id)
	asd.Close()

	responseLegajo(w, 202, legajo)

}


func LegajoAdd(w http.ResponseWriter, r *http.Request){
  
	decoder := json.NewDecoder(r.Body)

	var legajo_data Legajo
	//&nombre_var para decirle que es la var que no tiene datos y va a tener que rellenar
	err := decoder.Decode(&legajo_data)

	if(err != nil){
		panic(err)
	}

	//Para cerrar la lectura de algo
	defer r.Body.Close()

	log.Println(legajo_data)

	if err := db.Create(&legajo_data).Error; err != nil {
		fmt.Println("Error")
		fmt.Println(err)
	}

	responseLegajo(w, 202, legajo_data)

}


func LegajoUpdate(w http.ResponseWriter, r *http.Request){
 
	params := mux.Vars(r)
	legajo_id := params["id"]


	decoder := json.NewDecoder(r.Body)

	var legajo_data Legajo
	err := decoder.Decode(&legajo_data)

	if( err != nil ){
		panic(err)
		w.WriteHeader(500)
		return
	}
	fmt.Println(legajo_data)
	//cortar la lectura del body
	defer r.Body.Close()

	//Modifica el legajo que cumpla con la condición
	db.Model(Legajo{}).Where("id = ?", legajo_id).Updates(legajo_data)

	responseLegajo(w, 202, legajo_data)

}

func LegajoPatch(w http.ResponseWriter, r *http.Request){
 
	params := mux.Vars(r)
	legajo_id := params["id"]


	decoder := json.NewDecoder(r.Body)

	var legajo_data Legajo
	err := decoder.Decode(&legajo_data)

	if( err != nil ){
		panic(err)
		w.WriteHeader(500)
		return
	}
	fmt.Println(legajo_data)
	//cortar la lectura del body
	defer r.Body.Close()

	//Modifica el legajo que cumpla con la condición
	db.Model(Legajo{}).Where("id = ?", legajo_id).Updates(legajo_data)

	responseLegajo(w, 202, legajo_data)

}


type Message struct {
	Status string `json: "status"`
	Message string `json: "message"`
}
 //Forma de asociar el metodo con la estructura --> this puede ser cualquier nombre, no precisamente tiene que ser this
 //Va el * para pasarselo como puntero y quien use los metodos realmente modifiquen la estructura
func (this *Message) setStatus(data string){
	this.Status = data
}

 //Forma de asociar el metodo con la estructura --> this puede ser cualquier nombre, no precisamente tiene que ser this
 //Va el * para pasarselo como puntero y quien use los metodos realmente modifiquen la estructura
func (this *Message) setMessage(data string){
	this.Message = data
}


func LegajoRemove(w http.ResponseWriter, r *http.Request){

  	//Para obtener los parametros por la url
	params := mux.Vars(r)
	legajo_id := params["id"]

    //Eliminar legajo según condición
    //--Borrado Fisico
	fmt.Println(legajo_id)

    db.Unscoped().Where("id = ?", legajo_id).Delete(Legajo{})

    //--Borrado Logico
    //db.Where("descripcion = ?", "Probando Update").Delete(Legajo{})

    //db.Delete(Legajo{}, "descripcion = ?", "Probando Update")

    message := new(Message)
	message.setStatus("success")
	message.setMessage("El legajo con ID " +legajo_id+ " ha sido eliminado correctamente")

	results := message

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(results)
}

