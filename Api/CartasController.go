package api

import (
	http "net/http"
	services "restaurante-api/Services"

	"github.com/gin-gonic/gin"
)

func GetCartas(c *gin.Context) {
	cartas := services.GetCartas()

	if len(cartas) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No se encontraron cartas"})
		return
	}
	c.JSON(http.StatusOK, cartas)
}
