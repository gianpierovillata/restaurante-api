package domain

type Comida struct {
	ID           int      `json:"id"`
	Nombre       string   `json:"nombre"`
	Descripcion  string   `json:"descripcion"`
	Precio       float64  `json:"precio"`
	Ingredientes []string `json:"ingredientes"`
	Promocion    string   `json:"promocion"`
}
