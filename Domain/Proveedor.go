package domain

type Proveedor struct {
	ID         int    `json:"id"`
	Nombre     string `json:"nombre"`
	Telefono   string `json:"telefono"`
	Email      string `json:"email"`
	Direccion  string `json:"direccion"`
	TipoComida string `json:"tipo_comida"`
}
