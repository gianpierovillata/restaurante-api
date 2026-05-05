package main

import (
	"github.com/gin-gonic/gin"
)

type Promocion struct {
	ID          int     `json:"id"`
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Descuento   float64 `json:"descuento"`
	Codigo      string  `json:"codigo"`
}
type Comida struct {
	ID           int      `json:"id"`
	Nombre       string   `json:"nombre"`
	Descripcion  string   `json:"descripcion"`
	Precio       float64  `json:"precio"`
	Ingredientes []string `json:"ingredientes"`
	Promocion    string   `json:"promocion"`
}

var promociones = []Promocion{
	{ID: 1, Nombre: "Descuento de Verano", Descripcion: "10% de descuento en todas las comidas", Descuento: 10.0, Codigo: "VERANO"},
	{ID: 2, Nombre: "Promoción de Fin de Semana", Descripcion: "15% de descuento en comidas seleccionadas", Descuento: 15.0, Codigo: "FINDESEMANA"},
	{ID: 3, Nombre: "Oferta Especial", Descripcion: "20% de descuento en la comida del día", Descuento: 20.0, Codigo: "ESPECIAL"},
}
var comidas = []Comida{
	{ID: 1, Nombre: "Hamburguesa", Descripcion: "Hamburguesa con queso, lechuga y tomate", Precio: 8.99, Ingredientes: []string{"Carne de res", "Queso", "Lechuga", "Tomate", "Pan"}, Promocion: "VERANO"},
	{ID: 2, Nombre: "Pizza", Descripcion: "Pizza de pepperoni con extra queso", Precio: 12.99, Ingredientes: []string{"Masa de pizza", "Salsa de tomate", "Queso", "Pepperoni"}, Promocion: "FINDESEMANA"},
	{ID: 3, Nombre: "Ensalada César", Descripcion: "Ensalada César con pollo a la parrilla", Precio: 9.99, Ingredientes: []string{"Lechuga romana", "Pollo a la parrilla", "Aderezo César", "Crutones"}, Promocion: "ESPECIAL"},
}

func main() {

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Bienvenido a la api de restaurante"})

	})

	router.GET("/comidas", func(c *gin.Context) {

		c.JSON(200, comidas)
	})

	router.GET("/promociones", func(c *gin.Context) {
		if len(promociones) == 0 {
			c.JSON(404, gin.H{"message": "No se encontraron promociones"})
			return
		}

		c.JSON(200, promociones)
	})

	router.Run()
}
