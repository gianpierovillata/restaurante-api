package services

import (
	domain "restaurante-api/Domain"
	repository "restaurante-api/Repository"
)

func GetCartas() []domain.Carta {
	return repository.GetCartas()
}
