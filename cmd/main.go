package main

import (
	"encoding/json"
	"os"

	"github.com/Lelo88/GOWEB_EJERCICIO03/cmd/handler"
	"github.com/Lelo88/GOWEB_EJERCICIO03/internal/domain"
	"github.com/Lelo88/GOWEB_EJERCICIO03/internal/product"
	"github.com/gin-gonic/gin"
)

func main() {
	
	//carga de productos
	var listaProductos = []domain.Producto{}
	LoadProducts("products.json", &listaProductos)

	repo := product.NuevoRepositorio(listaProductos)
	service := product.NewService(repo)
	productHandler := handler.NewProductHandler(service)

    router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {ctx.String(200, "pong")})

	prods := router.Group("/productos")
	{
		prods.GET("/", productHandler.GetProductos())
		prods.GET("/id", productHandler.GetProductoById())
		prods.GET("/search", productHandler.PreciosMayores())
		prods.POST("/", productHandler.Create())
	}

	router.Run(":8080")
}

func LoadProducts(path string, listado *[]domain.Producto) {
	file, err := os.ReadFile(path)
    if err!= nil {
        panic(err)
    }
	
	err = json.Unmarshal(file, &listado)
    if err!= nil {
		panic(err)
	}
}