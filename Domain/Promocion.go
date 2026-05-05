package domain

type Promocion struct {
	ID          int     `json:"id"`
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Descuento   float64 `json:"descuento"`
	Codigo      string  `json:"codigo"`
}
