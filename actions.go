package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/xubiosueldos/autenticacion/publico"
	"github.com/xubiosueldos/conexionBD"
	"github.com/xubiosueldos/framework"
	"github.com/xubiosueldos/legajo/structLegajo"
)

func LegajoList(w http.ResponseWriter, r *http.Request) {

	tokenAutenticacion, tokenError := checkTokenValido(r)

	if tokenError != nil {
		errorToken(w, tokenError)
		return
	} else {

		db := obtenerDB(tokenAutenticacion)
		automigrateTablasPrivadas(db)
		defer db.Close()

		var legajos []structLegajo.Legajo

		//Lista todos los legajos
		db.Find(&legajos)

		framework.RespondJSON(w, http.StatusOK, legajos)
	}

}

func LegajoShow(w http.ResponseWriter, r *http.Request) {

	tokenAutenticacion, tokenError := checkTokenValido(r)

	if tokenError != nil {
		errorToken(w, tokenError)
		return
	} else {

		params := mux.Vars(r)
		legajo_id := params["id"]

		var legajo structLegajo.Legajo //Con &var --> lo que devuelve el metodo se le asigna a la var

		db := obtenerDB(tokenAutenticacion)
		automigrateTablasPrivadas(db)
		defer db.Close()

		//gorm:auto_preload se usa para que complete todos los struct con su informacion
		if err := db.Set("gorm:auto_preload", true).First(&legajo, "id = ?", legajo_id).Error; gorm.IsRecordNotFoundError(err) {
			framework.RespondError(w, http.StatusNotFound, err.Error())
			return
		}

		framework.RespondJSON(w, http.StatusOK, legajo)
	}

}

func LegajoAdd(w http.ResponseWriter, r *http.Request) {

	tokenAutenticacion, tokenError := checkTokenValido(r)

	if tokenError != nil {
		errorToken(w, tokenError)
		return
	} else {

		decoder := json.NewDecoder(r.Body)

		var legajo_data structLegajo.Legajo

		//&legajo_data para decirle que es la var que no tiene datos y va a tener que rellenar
		if err := decoder.Decode(&legajo_data); err != nil {
			framework.RespondError(w, http.StatusBadRequest, err.Error())
			return
		}

		defer r.Body.Close()

		db := obtenerDB(tokenAutenticacion)
		automigrateTablasPrivadas(db)
		defer db.Close()

		if err := db.Create(&legajo_data).Error; err != nil {
			framework.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		framework.RespondJSON(w, http.StatusCreated, legajo_data)
	}
}

func LegajoUpdate(w http.ResponseWriter, r *http.Request) {

	tokenAutenticacion, tokenError := checkTokenValido(r)

	if tokenError != nil {

		errorToken(w, tokenError)
		return
	} else {

		params := mux.Vars(r)
		//se convirtió el string en uint para poder comparar
		param_legajoid, _ := strconv.ParseUint(params["id"], 10, 64)
		p_legajoid := uint(param_legajoid)

		if p_legajoid == 0 {
			framework.RespondError(w, http.StatusNotFound, "Debe ingresar un ID en la url")
			return
		}

		decoder := json.NewDecoder(r.Body)

		var legajo_data structLegajo.Legajo

		if err := decoder.Decode(&legajo_data); err != nil {
			framework.RespondError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		legajoid := legajo_data.ID

		if p_legajoid == legajoid || legajoid == 0 {

			legajo_data.ID = p_legajoid

			db := obtenerDB(tokenAutenticacion)
			automigrateTablasPrivadas(db)
			defer db.Close()

			//abro una transacción para que si hay un error no persista en la DB
			tx := db.Begin()

			//modifico el legajo de acuerdo a lo enviado en el json
			if err := tx.Save(&legajo_data).Error; err != nil {
				tx.Rollback()
				framework.RespondError(w, http.StatusInternalServerError, err.Error())
				return
			}

			//despues de modificar, recorro los hijos asociados al legajo para ver si alguno fue eliminado logicamente y lo elimino de la BD
			if err := tx.Model(structLegajo.Hijo{}).Unscoped().Where("legajoid = ? AND deleted_at is not null", legajoid).Delete(structLegajo.Hijo{}).Error; err != nil {
				tx.Rollback()
				framework.RespondError(w, http.StatusInternalServerError, err.Error())
				return
			}

			//despues de modificar, recorro el conyuge asociado al legajo para ver si fue eliminado logicamente y lo elimino de la BD
			if err := tx.Model(structLegajo.Conyuge{}).Unscoped().Where("legajoid = ? AND deleted_at is not null", legajoid).Delete(structLegajo.Conyuge{}).Error; err != nil {
				tx.Rollback()
				framework.RespondError(w, http.StatusInternalServerError, err.Error())
				return
			}

			tx.Commit()

			framework.RespondJSON(w, http.StatusOK, legajo_data)

		} else {
			framework.RespondError(w, http.StatusNotFound, "El ID de la url debe ser el mismo que el del struct")
			return
		}
	}

}

func LegajoRemove(w http.ResponseWriter, r *http.Request) {

	tokenAutenticacion, tokenError := checkTokenValido(r)

	if tokenError != nil {

		errorToken(w, tokenError)
		return
	} else {

		//Para obtener los parametros por la url
		params := mux.Vars(r)
		legajo_id := params["id"]

		db := obtenerDB(tokenAutenticacion)
		automigrateTablasPrivadas(db)
		defer db.Close()

		//--Borrado Fisico
		if err := db.Unscoped().Where("id = ?", legajo_id).Delete(structLegajo.Legajo{}).Error; err != nil {

			framework.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		//--Borrado Logico
		//db.Where("descripcion = ?", "Probando Update").Delete(Legajo{})
		//db.Delete(Legajo{}, "descripcion = ?", "Probando Update")

		framework.RespondJSON(w, http.StatusOK, "El legajo con ID "+legajo_id+" ha sido eliminado correctamente")
	}

}

func obtenerDB(tokenAutenticacion *publico.TokenAutenticacion) *gorm.DB {

	token := *tokenAutenticacion
	tenant := token.Tenant

	return conexionBD.ConnectBD(tenant)

}

func automigrateTablasPrivadas(db *gorm.DB) {

	//para actualizar tablas...agrega columnas e indices, pero no elimina
	db.AutoMigrate(&structLegajo.Conyuge{}, &structLegajo.Hijo{}, &structLegajo.Legajo{})

	db.Model(&structLegajo.Hijo{}).AddForeignKey("legajoid", "legajo(id)", "CASCADE", "CASCADE")
	db.Model(&structLegajo.Conyuge{}).AddForeignKey("legajoid", "legajo(id)", "CASCADE", "CASCADE")
}

func errorToken(w http.ResponseWriter, tokenError *publico.Error) {
	errorToken := *tokenError
	framework.RespondError(w, errorToken.ErrorCodigo, errorToken.ErrorNombre)

}

func checkTokenValido(r *http.Request) (*publico.TokenAutenticacion, *publico.Error) {

	var tokenAutenticacion *publico.TokenAutenticacion
	var tokenError *publico.Error

	url := "http://localhost:8081/check-token"

	req, _ := http.NewRequest("GET", url, nil)

	header := r.Header.Get("Authorization")

	req.Header.Add("Authorization", header)

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
