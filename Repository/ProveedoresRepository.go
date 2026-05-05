package repository

import (
	domain "restaurante-api/Domain"
)

var proveedores = []domain.Proveedor{
	{ID: 1, Nombre: "Proveedor A", Telefono: "123456789", Email: "proveedorA@example.com", Direccion: "Calle 1", TipoComida: "Mexicana"},
	{ID: 2, Nombre: "Proveedor B", Telefono: "987654321", Email: "proveedorB@example.com", Direccion: "Calle 2", TipoComida: "Italiana"},
}

func GetProveedores() []domain.Proveedor {
	return proveedores
}
