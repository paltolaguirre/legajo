package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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
		return

	} else {
		token := *tokenAutenticacion
		tenant := token.Tenant

		db := conexionBD.ConnectBD(tenant)
		defer db.Close()

		var legajos []structLegajo.Legajo

		//Lista todos los legajos
		db.Find(&legajos)

		respondJSON(w, http.StatusOK, legajos)
	}

}

func LegajoShow(w http.ResponseWriter, r *http.Request) {

	tokenAutenticacion, tokenError := checkTokenValido(r)
	if tokenError != nil {
		errorToken := *tokenError
		respondError(w, errorToken.ErrorCodigo, errorToken.ErrorNombre)
		return

	} else {

		params := mux.Vars(r) //TODO: es global..? quizas usar el r
		legajo_id := params["id"]

		var legajo structLegajo.Legajo //Con &var --> lo que devuelve el metodo se le asigna a la var

		token := *tokenAutenticacion
		tenant := token.Tenant

		db := conexionBD.ConnectBD(tenant)
		defer db.Close()

		if err := db.Set("gorm:auto_preload", true).First(&legajo, "id = ?", legajo_id).Error; gorm.IsRecordNotFoundError(err) {
			respondError(w, http.StatusNotFound, err.Error())
			return
		}

		respondJSON(w, http.StatusOK, legajo)
	}

}

func LegajoAdd(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var legajo_data structLegajo.Legajo
	//&nombre_var para decirle que es la var que no tiene datos y va a tener que rellenar
	if err := decoder.Decode(&legajo_data); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	db := conexionBD.ConnectBD("tenant")
	defer db.Close()
	//Para cerrar la lectura de algo

	if err := db.Create(&legajo_data).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, legajo_data)

}

func LegajoUpdate(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	param_legajoid, _ := strconv.ParseUint(params["id"], 10, 64)
	p_legajoid := uint(param_legajoid)

	if p_legajoid == 0 {
		respondError(w, http.StatusNotFound, "Debe ingresar un ID en la url")
		return
	}

	decoder := json.NewDecoder(r.Body)

	var legajo_data structLegajo.Legajo

	if err := decoder.Decode(&legajo_data); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	legajoid := legajo_data.ID

	if p_legajoid == legajoid || legajoid == 0 {

		legajo_data.ID = uint(param_legajoid)

		db := conexionBD.ConnectBD("tenant")
		defer db.Close()

		//cortar la lectura del body

		//db.First(&legajo_data_db_current, "id = ?", param_legajoid)

		//Modifica el legajo que cumpla con la condición
		//db.Model(structLegajo.Legajo{}).Where("id = ?", legajo_id).Update(legajo_data)

		//db.Model(&legajo_data_db_current).Association("Hijos").Replace(legajo_data.Hijos)
		if err := db.Save(&legajo_data).Error; err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		//db.Model(structLegajo.Legajo{}.Hijos).Where("legajoid = ?", legajo_id).Update(legajo_data.Hijos)

		respondJSON(w, http.StatusOK, legajo_data)

	} else {
		respondError(w, http.StatusNotFound, "El ID de la url debe ser el mismo que el del struct")
		return
	}

}

/*func LegajoPatch(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	legajo_id := params["id"]

	decoder := json.NewDecoder(r.Body)

	var legajo_data structLegajo.Legajo
	err := decoder.Decode(&legajo_data)

	if err != nil {
		panic(err)
		respondError(w, 500, err.Error())
		return
	}
	fmt.Println(legajo_data)
	//cortar la lectura del body
	defer r.Body.Close()

	//Modifica el legajo que cumpla con la condición
	db.Model(structLegajo.Legajo{}).Where("id = ?", legajo_id).Updates(legajo_data)

	respondJSON(w, http.StatusOK, legajo_data)

}*/

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

	if err := db.Unscoped().Where("id = ?", legajo_id).Delete(structLegajo.Legajo{}).Error; err != nil {

		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	//--Borrado Logico
	//db.Where("descripcion = ?", "Probando Update").Delete(Legajo{})

	//db.Delete(Legajo{}, "descripcion = ?", "Probando Update")

	message := new(Message)
	message.setStatus("success")
	message.setMessage("El legajo con ID " + legajo_id + " ha sido eliminado correctamente")

	results := message

	respondJSON(w, http.StatusOK, results)

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
