package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/xubiosueldos/conexionBD"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/xubiosueldos/autenticacion/apiclientautenticacion"
	"github.com/xubiosueldos/conexionBD/Legajo/structLegajo"
	"github.com/xubiosueldos/framework"
	"github.com/xubiosueldos/monoliticComunication"
)

type IdsAEliminar struct {
	Ids []int `json:"ids"`
}

var nombreMicroservicio string = "legajo"

// Sirve para controlar si el server esta OK
func Healthy(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Healthy"))
}

func LegajoList(w http.ResponseWriter, r *http.Request) {

	tokenValido, tokenAutenticacion := apiclientautenticacion.CheckTokenValido(w, r)
	if tokenValido {

		tenant := apiclientautenticacion.ObtenerTenant(tokenAutenticacion)
		db := conexionBD.ObtenerDB(tenant)

		defer conexionBD.CerrarDB(db)
		var legajos []structLegajo.Legajo

		//Lista todos los legajos
		db.Find(&legajos)

		framework.RespondJSON(w, http.StatusOK, legajos)
	}

}

func LegajoShow(w http.ResponseWriter, r *http.Request) {

	tokenValido, tokenAutenticacion := apiclientautenticacion.CheckTokenValido(w, r)
	if tokenValido {
		params := mux.Vars(r)
		legajo_id := params["id"]

		var legajo structLegajo.Legajo //Con &var --> lo que devuelve el metodo se le asigna a la var

		tenant := apiclientautenticacion.ObtenerTenant(tokenAutenticacion)
		db := conexionBD.ObtenerDB(tenant)
		defer conexionBD.CerrarDB(db)

		//gorm:auto_preload se usa para que complete todos los struct con su informacion
		if err := db.Set("gorm:auto_preload", true).First(&legajo, "id = ?", legajo_id).Error; gorm.IsRecordNotFoundError(err) {
			framework.RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		centroCostoID := legajo.Centrodecostoid
		if centroCostoID != nil {
			centroDeCosto := monoliticComunication.Obtenercentrodecosto(w, r, tokenAutenticacion, strconv.Itoa(*centroCostoID))
			legajo.Centrodecosto = centroDeCosto
		}
		framework.RespondJSON(w, http.StatusOK, legajo)
	}

}

func LegajoAdd(w http.ResponseWriter, r *http.Request) {

	tokenValido, tokenAutenticacion := apiclientautenticacion.CheckTokenValido(w, r)
	if tokenValido {

		decoder := json.NewDecoder(r.Body)

		var legajo_data structLegajo.Legajo

		//&legajo_data para decirle que es la var que no tiene datos y va a tener que rellenar
		if err := decoder.Decode(&legajo_data); err != nil {
			framework.RespondError(w, http.StatusBadRequest, err.Error())
			return
		}

		defer r.Body.Close()

		tenant := apiclientautenticacion.ObtenerTenant(tokenAutenticacion)
		db := conexionBD.ObtenerDB(tenant)
		defer conexionBD.CerrarDB(db)

		centrodecostoid := legajo_data.Centrodecostoid
		if centrodecostoid != nil {
			if err_monolitico := monoliticComunication.Checkexistecentrodecosto(w, r, tokenAutenticacion, strconv.Itoa(*centrodecostoid)).Error; err_monolitico != nil {
				framework.RespondError(w, http.StatusInternalServerError, err_monolitico.Error())
				return
			}
		}

		if err := db.Create(&legajo_data).Error; err != nil {
			framework.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		framework.RespondJSON(w, http.StatusCreated, legajo_data)
	}
}

func LegajoUpdate(w http.ResponseWriter, r *http.Request) {

	tokenValido, tokenAutenticacion := apiclientautenticacion.CheckTokenValido(w, r)
	if tokenValido {

		params := mux.Vars(r)
		//se convirtió el string en uint para poder comparar
		param_legajoid, _ := strconv.ParseInt(params["id"], 10, 64)
		p_legajoid := int(param_legajoid)

		if p_legajoid == 0 {
			framework.RespondError(w, http.StatusNotFound, framework.IdParametroVacio)
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

			tenant := apiclientautenticacion.ObtenerTenant(tokenAutenticacion)
			db := conexionBD.ObtenerDB(tenant)
			defer conexionBD.CerrarDB(db)

			//abro una transacción para que si hay un error no persista en la DB
			tx := db.Begin()
			defer tx.Rollback()

			//modifico el legajo de acuerdo a lo enviado en el json
			centrodecostoid := legajo_data.Centrodecostoid
			if centrodecostoid != nil {
				if err := monoliticComunication.Checkexistecentrodecosto(w, r, tokenAutenticacion, strconv.Itoa(*centrodecostoid)).Error; err != nil {
					framework.RespondError(w, http.StatusInternalServerError, err.Error())
					return
				}
			}

			if err := tx.Save(&legajo_data).Error; err != nil {
				framework.RespondError(w, http.StatusInternalServerError, err.Error())
				return
			}

			//despues de modificar, recorro los hijos asociados al legajo para ver si alguno fue eliminado logicamente y lo elimino de la BD
			if err := tx.Model(structLegajo.Hijo{}).Unscoped().Where("legajoid = ? AND deleted_at is not null", legajoid).Delete(structLegajo.Hijo{}).Error; err != nil {
				framework.RespondError(w, http.StatusInternalServerError, err.Error())
				return
			}

			//despues de modificar, recorro el conyuge asociado al legajo para ver si fue eliminado logicamente y lo elimino de la BD
			if err := tx.Model(structLegajo.Conyuge{}).Unscoped().Where("legajoid = ? AND deleted_at is not null", legajoid).Delete(structLegajo.Conyuge{}).Error; err != nil {
				framework.RespondError(w, http.StatusInternalServerError, err.Error())
				return
			}

			tx.Commit()

			framework.RespondJSON(w, http.StatusOK, legajo_data)

		} else {
			framework.RespondError(w, http.StatusNotFound, framework.IdParametroDistintoStruct)
			return
		}
	}

}

func LegajoRemove(w http.ResponseWriter, r *http.Request) {

	tokenValido, tokenAutenticacion := apiclientautenticacion.CheckTokenValido(w, r)
	if tokenValido {
		//Para obtener los parametros por la url
		params := mux.Vars(r)
		legajo_id := params["id"]

		tenant := apiclientautenticacion.ObtenerTenant(tokenAutenticacion)
		db := conexionBD.ObtenerDB(tenant)

		defer conexionBD.CerrarDB(db)

		//--Borrado Fisico
		if err := db.Unscoped().Where("id = ?", legajo_id).Delete(structLegajo.Legajo{}).Error; err != nil {

			framework.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		//--Borrado Logico
		//db.Where("descripcion = ?", "Probando Update").Delete(Legajo{})
		//db.Delete(Legajo{}, "descripcion = ?", "Probando Update")

		framework.RespondJSON(w, http.StatusOK, framework.Legajo+legajo_id+framework.MicroservicioEliminado)
	}

}

func LegajosRemoveMasivo(w http.ResponseWriter, r *http.Request) {
	var resultadoDeEliminacion = make(map[int]string)
	tokenValido, tokenAutenticacion := apiclientautenticacion.CheckTokenValido(w, r)
	if tokenValido {

		var idsEliminar IdsAEliminar
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&idsEliminar); err != nil {
			framework.RespondError(w, http.StatusBadRequest, err.Error())
			return
		}

		tenant := apiclientautenticacion.ObtenerTenant(tokenAutenticacion)
		db := conexionBD.ObtenerDB(tenant)

		defer conexionBD.CerrarDB(db)

		if len(idsEliminar.Ids) > 0 {
			for i := 0; i < len(idsEliminar.Ids); i++ {
				legajo_id := idsEliminar.Ids[i]
				if err := db.Unscoped().Where("id = ?", legajo_id).Delete(structLegajo.Legajo{}).Error; err != nil {
					//framework.RespondError(w, http.StatusInternalServerError, err.Error())
					resultadoDeEliminacion[legajo_id] = string(err.Error())

				} else {
					resultadoDeEliminacion[legajo_id] = "Fue eliminado con exito"
				}
			}
		} else {
			framework.RespondError(w, http.StatusInternalServerError, "Seleccione por lo menos un registro")
		}

		framework.RespondJSON(w, http.StatusOK, resultadoDeEliminacion)
	}

}
