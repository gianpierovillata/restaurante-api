package repository

import domain "restaurante-api/Domain"

var comidas = []domain.Comida{
	{ID: 1, Nombre: "Hamburguesa", Descripcion: "Hamburguesa con queso, lechuga y tomate", Precio: 8.99, Ingredientes: []string{"Carne de res", "Queso", "Lechuga", "Tomate", "Pan"}, Promocion: "VERANO"},
	{ID: 2, Nombre: "Pizza", Descripcion: "Pizza de pepperoni con extra queso", Precio: 12.99, Ingredientes: []string{"Masa de pizza", "Salsa de tomate", "Queso", "Pepperoni"}, Promocion: "FINDESEMANA"},
	{ID: 3, Nombre: "Ensalada César", Descripcion: "Ensalada César con pollo a la parrilla", Precio: 9.99, Ingredientes: []string{"Lechuga romana", "Pollo a la parrilla", "Aderezo César", "Crutones"}, Promocion: "ESPECIAL"},
}

func GetComidas() []domain.Comida {
	return comidas
}
