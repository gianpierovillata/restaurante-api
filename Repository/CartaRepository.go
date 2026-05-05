package repository

import (
	domain "restaurante-api/Domain"
)

var cartas = []domain.Carta{
	{ID: 1, Nombre: "Carta de Verano", Fecha: "2024-06-01", Menu: []domain.Comida{
		{ID: 1, Nombre: "Ensalada César", Descripcion: "Lechuga, pollo, crutones y aderezo César", Precio: 8.99},
		{ID: 2, Nombre: "Gazpacho", Descripcion: "Sopa fría de tomate y verduras", Precio: 6.99},
	}},
	{ID: 2, Nombre: "Carta de Invierno", Fecha: "2024-12-01", Menu: []domain.Comida{
		{ID: 3, Nombre: "Sopa de Calabaza", Descripcion: "Sopa caliente de calabaza con especias", Precio: 7.99},
		{ID: 4, Nombre: "Estofado de Carne", Descripcion: "Carne guisada con verduras y salsa", Precio: 14.99},
	}},
}

func GetCartas() []domain.Carta {
	return cartas
}
