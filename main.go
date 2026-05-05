package main

import (
	api "restaurante-api/Api"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Bienvenido a la api de restaurante"})

	})

	router.GET("/comidas", api.GetComidas)

	router.GET("/proveedores", api.GetProveedores)

	router.GET("/cartas", api.GetCartas)

	router.Run(":8081")
}
