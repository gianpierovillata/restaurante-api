package api

import (
	http "net/http"
	services "restaurante-api/Services"

	"github.com/gin-gonic/gin"
)

func GetProveedores(c *gin.Context) {
	proveedores := services.GetProveedores()

	if len(proveedores) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No se encontraron proveedores"})
		return
	}
	c.JSON(http.StatusOK, proveedores)
}
