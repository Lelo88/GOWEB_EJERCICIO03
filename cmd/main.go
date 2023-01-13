package main

import (
	"github.com/Lelo88/GOWEB_EJERCICIO03/cmd/handler"
	"github.com/Lelo88/GOWEB_EJERCICIO03/product"
	"github.com/gin-gonic/gin"
)

func main() {
	
	//carga de productos
	product.LoadProducts("products.json")
	
	//inicio el server
	router := gin.Default()
	
	//1. detallo un acceso ping 
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	
	//2.detallo el acceso a los productos
	products := router.Group("/products")
	{
		products.GET("", handler.GetProductos()) //localhost:8080/products --> si colocamos mal algo, nos va a tirar el error indicado
		products.GET("/:id", handler.GetProductoById()) //localhost:8080/products/numero entero --> si colocamos mal algo, nos va a tirar el error indicado
		products.GET("/search", handler.PreciosMayores()) //localhost:8080/products/search?price=900 --> si colocamos mal algo, nos va a tirar el error indicado
		products.POST("", handler.CreateProduct()) //2. llamamos al metodo post para crear un producto, en caso de algun dato erroneo, evalua el error
		//2.añadimos el  metodo POST para la segun parte del ejercicio
	}	

    router.Run()
}