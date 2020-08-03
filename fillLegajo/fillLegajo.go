package fillLegajo

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/Legajo/structLegajo"
)

func CheckAndFill(legajo *structLegajo.Legajo, db * gorm.DB) error {

	var err error

	if legajo.Legajo == "" {
		return errors.New("El numero de legajo es obligatorio")
	}

	if existeLegajo(legajo.ID, legajo.Legajo, db) {
		return errors.New("Ya existe un legajo con ese mismo numero de legajo")
	}

	if existeCuil(legajo.ID, legajo.Cuil, db) {
		return errors.New("Ya existe un legajo con ese mismo cuil")
	}

	if legajo.Cuil == nil || *legajo.Cuil == ""{
		return errors.New("El numero de cuil es obligatorio")
	}

	if legajo.Apellido == nil || *legajo.Apellido == ""{
		return errors.New("El apellido es obligatorio")
	}

	if legajo.Nombre == nil || *legajo.Nombre == ""{
		return errors.New("El nombre es obligatorio")
	}

	if legajo.Nombre == nil || *legajo.Nombre == ""{
		return errors.New("El nombre es obligatorio")
	}

	if legajo.Horasmensualesnormales == "" {
		return errors.New("Las horas mensuales normales son obligatorias")
	}

	if legajo.Categoria == nil || *legajo.Categoria == ""{
		return errors.New("La categoria es obligatoria")
	}

	if legajo.Tarea == nil || *legajo.Tarea == ""{
		return errors.New("La tarea es obligatoria")
	}

	if legajo.Condicionid == nil {
		legajo.Condicionid, err = findIdCondicion(legajo.Condicion, db)
		legajo.Condicion = nil

		if err != nil {
			return err
		}

	}

	if legajo.Situacionid == nil {
		legajo.Situacionid, err = findIdSituacion(legajo.Situacion, db)
		legajo.Situacion = nil

		if err != nil {
			return err
		}
	}

	if legajo.Condicionsiniestradoid == nil {
		legajo.Condicionsiniestradoid, err = findIdCondicionSiniestrado(legajo.Condicionsiniestrado, db)
		legajo.Condicionsiniestrado = nil

		if err != nil {
			return err
		}
	}

	if legajo.Estadocivilid == nil {
		legajo.Estadocivilid, err = findIdEstadoCivil(legajo.Estadocivil, db)
		legajo.Estadocivil = nil
		if err != nil {
			return err
		}
	}

	if legajo.Localidadid == nil {
		legajo.Localidadid, err = findIdLocalidad(legajo.Localidad, db)
		legajo.Localidad = nil
		if err != nil {
			return err
		}
	}

	if legajo.Modalidadcontratacionid == nil {
		legajo.Modalidadcontratacionid, err = findIdModalidadcontratacion(legajo.Modalidadcontratacion, db)
		legajo.Modalidadcontratacion = nil
		if err != nil {
			return err
		}
	}

	if legajo.Obrasocialid == nil {
		legajo.Obrasocialid, err = findIdObraSocial(legajo.Obrasocial, db)
		legajo.Obrasocial = nil
		if err != nil {
			return err
		}
	}

	if legajo.Paisid == nil {
      		legajo.Paisid, err = findIdPais(legajo.Pais, db)
      		legajo.Pais = nil
		if err != nil {
			return err
		}
	}

	if legajo.Provinciaid == nil {
		legajo.Provinciaid, err = findIdProvincia(legajo.Provincia, db)
		legajo.Provincia = nil
		if err != nil {
			return err
		}
	}

	if legajo.Hijos != nil {
		for _, hijo := range legajo.Hijos {
			if hijo.Nombre == "" {
				return errors.New("El nombre de los hijos es obligatorio")
			}
			if hijo.Apellido == "" {
				return errors.New("El apellido de los hijos es obligatorio")
			}
			if hijo.Cuil == "" {
				return errors.New("El cuil de los hijos es obligatorio")
			}
			if hijo.Obrasocialid == nil {
				hijo.Obrasocialid, err = findIdObraSocial(hijo.Obrasocial, db)
				hijo.Obrasocial = nil
				if err != nil {
					return errors.New("El hijo " + hijo.Nombre + " " + hijo.Apellido + " no se pudo crear: " + err.Error())
				}
			}
		}
	}

	if legajo.Conyuge != nil {
		for _, conyuge := range legajo.Conyuge {
			if conyuge.Nombre == "" {
				return errors.New("El nombre del conyuge es obligatorio")
			}
			if conyuge.Apellido == "" {
				return errors.New("El apellido del conyuge es obligatorio")
			}
			if conyuge.Cuil == "" {
				return errors.New("El cuil del conyuge es obligatorio")
			}
			if conyuge.Obrasocialid == nil {
				conyuge.Obrasocialid, err = findIdObraSocial(conyuge.Obrasocial, db)
				conyuge.Obrasocial = nil
				if err != nil {
					return errors.New("El conyuge " + conyuge.Nombre + " " + conyuge.Apellido + " no se pudo crear: " + err.Error())
				}
			}
		}
	}

	return nil

}

func existeCuil(id int, cuil *string, db *gorm.DB) bool {
	var count int

	db.Model(&structLegajo.Legajo{}).Where("cuil = ? AND id != ?", cuil, id).Count(&count)

	return count > 0
}

func existeLegajo(id int, legajo string, db *gorm.DB) bool {
	var count int

	db.Model(&structLegajo.Legajo{}).Where("legajo = ? AND id != ?", legajo, id).Count(&count)

	return count > 0
}

func findIdProvincia(provincia *structLegajo.Provincia, db *gorm.DB) (*int, error) {
	if provincia == nil {
		return nil, errors.New("La provincia es obligatoria")
	}
	if provincia.Codigo != ""{
		if db.Where("codigo ilike ?", provincia.Codigo).First(&provincia).RecordNotFound() {
			return nil, errors.New("No existe la provincia con el codigo " + provincia.Codigo)
		}
	} else if provincia.Nombre != "" {
		if db.Where("nombre ilike ?", provincia.Nombre).First(&provincia).RecordNotFound(){
			return nil, errors.New("No existe la provincia con el nombre " + provincia.Nombre)
		}
	} else {
		return nil, errors.New("La provincia es obligatoria")
	}

	return &provincia.ID, nil
}

func findIdPais(pais *structLegajo.Pais, db *gorm.DB) (*int, error) {
	if pais.Codigo != ""{
		if db.Where("codigo ilike ?", pais.Codigo).First(&pais).RecordNotFound() {
			return nil, errors.New("No existe el pais con el codigo " + pais.Codigo)
		}
	} else if pais.Nombre != "" {
		if db.Where("nombre ilike ?", pais.Nombre).First(&pais).RecordNotFound(){
			return nil, errors.New("No existe el pais con el nombre " + pais.Nombre)
		}
	} else {
		return nil, errors.New("El pais es obligatorio")
	}

	return &pais.ID, nil
}

func findIdObraSocial(obrasocial *structLegajo.Obrasocial, db *gorm.DB) (*int, error) {
	if obrasocial == nil {
		return nil, errors.New("La obra social es obligatoria")
	}
	if obrasocial.Codigo != ""{
		if db.Where("codigo ilike ?", obrasocial.Codigo).First(&obrasocial).RecordNotFound() {
			return nil, errors.New("No existe la obra social con el codigo " + obrasocial.Codigo)
		}
	} else if obrasocial.Nombre != "" {
		if db.Where("nombre ilike ?", obrasocial.Nombre).First(&obrasocial).RecordNotFound(){
			return nil, errors.New("No existe la obra social con el nombre " + obrasocial.Nombre)
		}
	} else {
		return nil, errors.New("La obra social es obligatoria")
	}

	return &obrasocial.ID, nil
}

func findIdModalidadcontratacion(modalidadcontratacion *structLegajo.Modalidadcontratacion, db *gorm.DB) (*int, error) {
	if modalidadcontratacion == nil {
		return nil, errors.New("La modalidad de contratacion es obligatoria")
	}

	if modalidadcontratacion.Codigo != ""{
		if db.Where("codigo ilike ?", modalidadcontratacion.Codigo).First(&modalidadcontratacion).RecordNotFound() {
			return nil, errors.New("No existe la modalidad de contratacion con el codigo " + modalidadcontratacion.Codigo)
		}
	} else if modalidadcontratacion.Nombre != "" {
		if db.Where("nombre ilike ?", modalidadcontratacion.Nombre).First(&modalidadcontratacion).RecordNotFound(){
			return nil, errors.New("No existe la modalidad de contratacion con el nombre " + modalidadcontratacion.Nombre)
		}
	} else {
		return nil, errors.New("La modalidad de contratacion es obligatoria")
	}

	return &modalidadcontratacion.ID, nil
}

func findIdLocalidad(localidad *structLegajo.Localidad, db *gorm.DB) (*int,error) {
	if localidad == nil {
		return nil, errors.New("La localidad es obligatoria")
	}

	if localidad.Codigo != ""{
		if db.Where("codigo ilike ?", localidad.Codigo).First(&localidad).RecordNotFound() {
			return nil, errors.New("No existe la localidad con el codigo " + localidad.Codigo)
		}
	} else if localidad.Nombre != "" {
		if db.Where("nombre ilike ?", localidad.Nombre).First(&localidad).RecordNotFound(){
			return nil, errors.New("No existe la localidad con el nombre " + localidad.Nombre)
		}
	} else {
		return nil, errors.New("La localidad es obligatoria")
	}

	return &localidad.ID, nil
}

func findIdEstadoCivil(estadocivil *structLegajo.Estadocivil, db *gorm.DB) (*int, error) {
	if estadocivil == nil {
		return nil, errors.New("El estado civil es obligatorio")
	}

	if estadocivil.Codigo != ""{
		if db.Where("codigo ilike ?", estadocivil.Codigo).First(&estadocivil).RecordNotFound() {
			return nil, errors.New("No existe el estado civil con el codigo " + estadocivil.Codigo)
		}
	} else if estadocivil.Nombre != "" {
		if db.Where("nombre ilike ?", estadocivil.Nombre).First(&estadocivil).RecordNotFound(){
			return nil, errors.New("No existe el estado civil con el nombre " + estadocivil.Nombre)
		}
	} else {
		return nil, errors.New("El estado civil es obligatorio")
	}

	return &estadocivil.ID, nil
}

func findIdCondicionSiniestrado(condicionsiniestrado *structLegajo.Condicionsiniestrado, db *gorm.DB) (*int, error) {

	if condicionsiniestrado == nil {
		return nil, errors.New("La condicion de siniestrado es obligatoria")
	}

	if condicionsiniestrado.Codigo != ""{
		if db.Where("codigo ilike ?", condicionsiniestrado.Codigo).First(&condicionsiniestrado).RecordNotFound() {
			return nil, errors.New("No existe la condicion de siniestrado con el codigo " + condicionsiniestrado.Codigo)
		}
	} else if condicionsiniestrado.Nombre != "" {
		if db.Where("nombre ilike ?", condicionsiniestrado.Nombre).First(&condicionsiniestrado).RecordNotFound(){
			return nil, errors.New("No existe la condicion de siniestrado con el nombre " + condicionsiniestrado.Nombre)
		}
	} else {
		return nil, errors.New("La condicion de siniestrado es obligatoria")
	}

	return &condicionsiniestrado.ID, nil
}

func findIdSituacion(situacion *structLegajo.Situacion, db *gorm.DB) (*int, error) {
	if situacion == nil {
		return nil, errors.New("La situacion es obligatoria")
	}
	if situacion.Codigo != ""{
		if db.Where("codigo ilike ?", situacion.Codigo).First(&situacion).RecordNotFound() {
			return nil, errors.New("No existe la situacion con el codigo " + situacion.Codigo)
		}
	} else if situacion.Nombre != "" {
		if db.Where("nombre ilike ?", situacion.Nombre).First(&situacion).RecordNotFound(){
			return nil, errors.New("No existe la situacion con el nombre " + situacion.Nombre)
		}
	} else {
		return nil, errors.New("La situacion es obligatoria")
	}

	return &situacion.ID, nil
}

func findIdCondicion(condicion *structLegajo.Condicion, db *gorm.DB) (*int,error) {
	if condicion == nil {
		return nil, errors.New("La condicion es obligatoria")
	}
	if condicion.Codigo != ""{
		if db.Where("codigo ilike ?", condicion.Codigo).First(&condicion).RecordNotFound() {
			return nil, errors.New("No existe la condicion con el codigo " + condicion.Codigo)
		}
	} else if condicion.Nombre != "" {
		if db.Where("nombre ilike ?", condicion.Nombre).First(&condicion).RecordNotFound(){
			return nil, errors.New("No existe la condicion con el nombre " + condicion.Nombre)
		}
	} else {
		return nil, errors.New("La condicion es obligatoria")
	}

	return &condicion.ID, nil

}
