package services

import (
	domain "restaurante-api/Domain"
	repository "restaurante-api/Repository"
)

func GetComidas() []domain.Comida {
	return repository.GetComidas()
}
