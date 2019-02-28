package structLegajo


type Legajo struct {
	gorm.Model
	Nombre string
	Codigo string
	Descripcion string
	Activo int
}