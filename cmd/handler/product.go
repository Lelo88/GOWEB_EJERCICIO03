package handler

import (
	"strconv"

	"github.com/Lelo88/GOWEB_EJERCICIO03/product"
	"github.com/gin-gonic/gin"
)

func GetProductos() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, product.Productos)
	}
}

//funcion que me va a devolver un producto por ID. Si no existe, return 404
func GetProductoById() gin.HandlerFunc{
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid id"})
			return
		}
		for _, product := range product.Productos {
			if product.Id == id {
				c.JSON(200, product)
				return
			}
		}
		c.JSON(404, gin.H{"error": "product not found"})
	}
}	

func PreciosMayores()gin.HandlerFunc{
	return func(c *gin.Context){
		var resultado []product.Producto

		precio, err := strconv.ParseFloat(c.Query("price"), 64)
		
		if err!= nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _,p := range product.Productos {
		if p.Price > precio{
			resultado = append(resultado, p)
		}
	}
	
	c.JSON(200, resultado)
	}	
}

func CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
        
		//si lo que ingresamos no es un objeto json, nos devuelve un error
		var producto product.Producto
		
		err := c.ShouldBindJSON(&producto)
        if err!= nil {
            c.JSON(400, gin.H{"error": "producto inválido"})
			return
		}

		//si lo que ingresamos no es valido, nos devuelve un error
		valido, err := product.ValidarProducto(&producto)
		if err!=nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return 
		}
		
		valido, err = product.ValidarFecha(&producto)
		if !valido{
			c.JSON(400, gin.H{"error": err.Error()})
			return 
		}
		
		valido = product.ValidaCodigoID(producto.Code_Value)
		if !valido {
			c.JSON(400, gin.H{"error": "codigo existente"})
            return
        }

		producto.Id = len(product.Productos) + 1
		product.Productos = append(product.Productos, producto)
		c.JSON(201, producto)
	}
}