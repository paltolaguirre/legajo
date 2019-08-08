package structLegajo

type Centrodecosto struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Codigo      string `json:"codigo"`
	Descripcion string `json:"descripcion"`
	Activo      int    `json:"activo"`
}
