package main

import (
	"encoding/json"
	"strconv"
	"os"
	"github.com/gin-gonic/gin"
)

//creo la estructura para traerme los datos del archivo json
type Producto struct {
	Id 				int		`json:"id"`
	Name 			string 	`json:"name"`
	Quantity 		int 	`json:"quantity"`
	Code_Value 		string 	`json:"code_value"`
	Is_Published 	bool 	`json:"is_published"`
	Expiration 		string 	`json:"expiration"`
	Price 			float64 `json:"price"`
}

//creo un slice vacio para almacenar los datos del archivo json en memoria
var productos = []Producto{}

//funcion que va a leer el archivo json y los va a llevar al slice vacio
func loadProducts(path string, list *[]Producto) {
	file, err := os.ReadFile(path)
    if err!= nil {
        panic(err)
    }

	err = json.Unmarshal(file, list)
    if err!= nil {
		panic(err)
	}
}

//funcion que me va a traer todos los productos del listado en memoria
func getProductos() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, productos)
	}
}

//funcion que me va a devolver un producto por ID
func getProductoById() gin.HandlerFunc{
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid id"})
			return
		}
		for _, product := range productos {
			if product.Id == id {
				c.JSON(200, product)
				return
			}
		}
		c.JSON(404, gin.H{"error": "product not found"})
	}
}	

func preciosMayores()gin.HandlerFunc{
	return func(c *gin.Context){
		var resultado []Producto

		precio, err := strconv.ParseFloat(c.Query("price"), 64)
		
		if err!= nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _,p := range productos {
		if p.Price > precio{
			resultado = append(resultado, p)
		}
	}
	
	c.JSON(200, resultado)
	}	
}


func main() {
	
	//carga de productos
	loadProducts("products.json", &productos)
	
	//inicio el server
	router := gin.Default()
	
	//1. detallo un acceso ping 
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	
	//2.detallo el acceso para los productos
	products := router.Group("/products")
	{
		products.GET("", getProductos())
		products.GET("/:id", getProductoById())
		products.GET("/search", preciosMayores())
    }	
	
	

    router.Run()
}