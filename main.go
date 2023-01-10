package main

import (
	"encoding/json"
	"strconv"
	//"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id 				int		`json:"id"`
	Name 			string 	`json:"name"`
	Quantity 		int 	`json:"quantity"`
	Code_Value 		string 	`json:"code_value"`
	Is_Published 	bool 	`json:"is_published"`
	Expiration 		string 	`json:"expiration"`
	Price 			float64 `json:"price"`
}

var productos []Producto


func getProductos() []Producto {
	archivoJson, err := os.Open("products.json")
	if err!=nil {
		log.Fatal(err)
	}

	defer archivoJson.Close()

	bytes, err := io.ReadAll(archivoJson)
	if err!= nil { 
		log.Fatal(err)
    }

	json.Unmarshal([]byte(bytes), &productos)

	return productos
}

func getProductoById(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	
	if err!=nil {
		c.String(404, "Empleado %s no encontrado", c.Param("id"))
    }
	
	for _,producto := range productos {
		if producto.Id == id{
            c.JSON(200, producto)
            return
        }
    }
}	


func main() {
	router := gin.Default()
	

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.GET("/products", func(c *gin.Context) {
		listado := getProductos()
		c.JSON(200, listado)
	})

	router.GET("/products/:id", getProductoById)

    router.Run()
}