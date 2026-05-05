package domain

type Carta struct {
	ID     int      `json:"id"`
	Nombre string   `json:"nombre"`
	Fecha  string   `json:"fecha"`
	Menu   []Comida `json:"menu"`
	
}
