package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/xubiosueldos/autenticacion/publico"
	"github.com/xubiosueldos/conexionBD"
	"github.com/xubiosueldos/legajo/structLegajo"
)

var db *gorm.DB
var err error

func respondJSON(w http.ResponseWriter, status int, results interface{}) {

	response, err := json.Marshal(results)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))

}

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}

func LegajoList(w http.ResponseWriter, r *http.Request) {

	tokenAutenticacion, tokenError := checkTokenValido(r)

	if tokenError != nil {

		errorToken := *tokenError

		respondError(w, errorToken.ErrorCodigo, errorToken.ErrorNombre)

	} else {
		token := *tokenAutenticacion
		tenant := token.Tenant
		fmt.Println(tenant)
		db := conexionBD.ConnectBD(tenant)

		var legajos []structLegajo.Legajo

		//Lista todos los legajos
		fmt.Println("Los legajos de la BD son: ")
		db.Find(&legajos)

		fmt.Println(legajos)
		respondJSON(w, 202, legajos)
	}

}

func LegajoShow(w http.ResponseWriter, r *http.Request) {

	tokenAutenticacion, tokenError := checkTokenValido(r)
	if tokenError != nil {
		errorToken := *tokenError
		respondError(w, errorToken.ErrorCodigo, errorToken.ErrorNombre)
		fmt.Println(errorToken)

	} else {

		params := mux.Vars(r) //TODO: es global..? quizas usar el r
		legajo_id := params["id"]
		fmt.Println(legajo_id)

		var legajo structLegajo.Legajo //Con &var --> lo que devuelve el metodo se le asigna a la var

		//db.First(&legajo, "id = ?", legajo_id)
		token := *tokenAutenticacion
		tenant := token.Tenant
		db := conexionBD.ConnectBD(tenant)
		db.Set("gorm:auto_preload", true).First(&legajo, "id = ?", legajo_id)
		db.Close()

		respondJSON(w, 202, legajo)
	}

}

func LegajoAdd(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var legajo_data structLegajo.Legajo
	//&nombre_var para decirle que es la var que no tiene datos y va a tener que rellenar
	err := decoder.Decode(&legajo_data)

	if err != nil {
		panic(err)
	}
	db := conexionBD.ConnectBD("tenant")

	//Para cerrar la lectura de algo
	defer r.Body.Close()

	log.Println(legajo_data)

	if err := db.Create(&legajo_data).Error; err != nil {
		fmt.Println("Error")
		fmt.Println(err)
	}

	respondJSON(w, 202, legajo_data)

}

func LegajoUpdate(w http.ResponseWriter, r *http.Request) {

	//params := mux.Vars(r)
	//legajo_id := params["id"]

	decoder := json.NewDecoder(r.Body)

	var legajo, legajo_data structLegajo.Legajo
	err := decoder.Decode(&legajo_data)

	if err != nil {
		panic(err)
		w.WriteHeader(500)
		return
	}
	fmt.Println(legajo_data)
	db := conexionBD.ConnectBD("tenant")

	//cortar la lectura del body
	defer r.Body.Close()

	//Modifica el legajo que cumpla con la condición
	//db.Model(structLegajo.Legajo{}).Where("id = ?", legajo_id).Update(legajo_data)
	db.Model(&legajo).Association("Hijos").Replace(legajo_data)
	db.Save(&legajo_data)

	//db.Model(structLegajo.Legajo{}.Hijos).Where("legajoid = ?", legajo_id).Update(legajo_data.Hijos)

	respondJSON(w, 202, legajo_data)

}

func LegajoPatch(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	legajo_id := params["id"]

	decoder := json.NewDecoder(r.Body)

	var legajo_data structLegajo.Legajo
	err := decoder.Decode(&legajo_data)

	if err != nil {
		panic(err)
		w.WriteHeader(500)
		return
	}
	fmt.Println(legajo_data)
	//cortar la lectura del body
	defer r.Body.Close()

	//Modifica el legajo que cumpla con la condición
	db.Model(structLegajo.Legajo{}).Where("id = ?", legajo_id).Updates(legajo_data)

	respondJSON(w, 202, legajo_data)

}

type Message struct {
	Status  string `json: "status"`
	Message string `json: "message"`
}

//Forma de asociar el metodo con la estructura --> this puede ser cualquier nombre, no precisamente tiene que ser this
//Va el * para pasarselo como puntero y quien use los metodos realmente modifiquen la estructura
func (this *Message) setStatus(data string) {
	this.Status = data
}

//Forma de asociar el metodo con la estructura --> this puede ser cualquier nombre, no precisamente tiene que ser this
//Va el * para pasarselo como puntero y quien use los metodos realmente modifiquen la estructura
func (this *Message) setMessage(data string) {
	this.Message = data
}

func LegajoRemove(w http.ResponseWriter, r *http.Request) {

	//Para obtener los parametros por la url
	params := mux.Vars(r)
	legajo_id := params["id"]

	//Eliminar legajo según condición
	//--Borrado Fisico
	fmt.Println(legajo_id)
	db := conexionBD.ConnectBD("tenant")

	db.Unscoped().Where("id = ?", legajo_id).Delete(structLegajo.Legajo{})

	//--Borrado Logico
	//db.Where("descripcion = ?", "Probando Update").Delete(Legajo{})

	//db.Delete(Legajo{}, "descripcion = ?", "Probando Update")

	message := new(Message)
	message.setStatus("success")
	message.setMessage("El legajo con ID " + legajo_id + " ha sido eliminado correctamente")

	results := message

	respondJSON(w, 200, results)

}

func checkTokenValido(r *http.Request) (*publico.TokenAutenticacion, *publico.Error) {

	var tokenAutenticacion *publico.TokenAutenticacion
	var tokenError *publico.Error

	url := "http://localhost:8081/check-token"

	req, _ := http.NewRequest("GET", url, nil)

	token := r.Header.Get("Token")

	req.Header.Add("token", token)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != 400 {

		// tokenAutenticacion = &(TokenAutenticacion{})
		tokenAutenticacion = new(publico.TokenAutenticacion)
		json.Unmarshal([]byte(string(body)), tokenAutenticacion)

	} else {
		tokenError = new(publico.Error)
		json.Unmarshal([]byte(string(body)), tokenError)

	}

	return tokenAutenticacion, tokenError
}

func ProvinciaShow(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r) //TODO: es global..? quizas usar el r
	legajo_id := params["id"]
	fmt.Println(legajo_id)

	var provincia structLegajo.Provincia //Con &var --> lo que devuelve el metodo se le asigna a la var

	//db.First(&legajo, "id = ?", legajo_id)

	db := conexionBD.ConnectBD("")
	//db.First(&provincia, "id = ?", legajo_id)

	/*	var p structLegajo.Provincia

		p.Nombre = "Buenos Aires"
		db.Create(&p)*/

	/*var p structLegajo.Pais

	p.Nombre = "Argentina"
	db.Create(&p)
	*/
	/*var p structLegajo.Provincia
	p.PaisId = 2
	db.Model(structLegajo.Provincia{}).Where("id = ?", 2).Updates(p)
	db.Where(&p).First(&p)*/

	//db.First(&provincia, legajo_id)

	db.Set("gorm:auto_preload", true).Where("id = ?", legajo_id).Find(&provincia)

	//db.Find(&provincia, structLegajo.Provincia{Nombre: "Buenos Aires"})

	db.Close()

	respondJSON(w, 202, provincia)
}
