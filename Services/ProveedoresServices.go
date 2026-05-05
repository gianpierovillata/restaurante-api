package services

import (
	domain "restaurante-api/Domain"
	repository "restaurante-api/Repository"
)

func GetProveedores() []domain.Proveedor {
	return repository.GetProveedores()
}
