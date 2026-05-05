package main

import (
	services "restaurante-api/Services"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Bienvenido a la api de restaurante"})

	})

	router.GET("/comidas", func(c *gin.Context) {

		c.JSON(200, services.GetComidas())
	})

	router.GET("/proveedores", func(c *gin.Context) {

		c.JSON(200, services.GetProveedores())
	})

	router.GET("/cartas", func(c *gin.Context) {

		c.JSON(200, services.GetCartas())
	})

	router.Run()
}
