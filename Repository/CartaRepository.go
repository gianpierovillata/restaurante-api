package repository

import (
	domain "restaurante-api/Domain"
)

var cartas = []domain.Carta{
	{ID: 1, Nombre: "Carta de Verano", Fecha: "2024-06-01", Menu: []domain.Comida{
		{ID: 1, Nombre: "Ensalada César", Descripcion: "Lechuga, pollo, crutones y aderezo César", Precio: 8.99, Ingredientes: []string{"Lechuga", "Pollo", "Crutones", "Aderezo César"}, Promocion: "VERANO"},
		{ID: 2, Nombre: "Gazpacho", Descripcion: "Sopa fría de tomate y verduras", Precio: 6.99, Ingredientes: []string{"Tomate", "Pepino", "Pimiento", "Ajo", "Aceite de oliva"}, Promocion: "VERANO"},
	}},
	{ID: 2, Nombre: "Carta de Invierno", Fecha: "2024-12-01", Menu: []domain.Comida{
		{ID: 3, Nombre: "Sopa de Calabaza", Descripcion: "Sopa caliente de calabaza con especias", Precio: 7.99, Ingredientes: []string{"Calabaza", "Especias"}, Promocion: "INVIERNO"},
		{ID: 4, Nombre: "Estofado de Res", Descripcion: "Carne de res cocida lentamente con verduras", Precio: 12.99, Ingredientes: []string{"Carne", "Verduras", "Salsa"}, Promocion: "INVIERNO"},
	}},
}

func GetCartas() []domain.Carta {
	return cartas
}
