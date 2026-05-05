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

	/* router.GET("/promociones", func(c *gin.Context) {
		if len(services.GetPromociones()) == 0 {
			c.JSON(404, gin.H{"message": "No se encontraron promociones"})
			return
		}

		c.JSON(200, services.GetPromociones())
	}) */

	router.Run()
}
