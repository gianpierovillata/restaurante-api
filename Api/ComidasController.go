package api

import (
	http "net/http"
	services "restaurante-api/Services"

	"github.com/gin-gonic/gin"
)

func GetComidas(c *gin.Context) {
	comidas := services.GetComidas()

	if len(comidas) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No se encontraron comidas"})
		return
	}
	c.JSON(http.StatusOK, comidas)
}
