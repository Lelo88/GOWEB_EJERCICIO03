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

	repo := product.NuevoRepositorio(listaProductos) //creo un repositorio nuevo con los datos provenientes del archivo products.json
	service := product.NewService(repo) //creo los servicios correspondientes a repo, con todos sus metodos
	productHandler := handler.NewProductHandler(service) //el handler que creamos correspondera al servicio que creemos

    router := gin.Default() //configuramos el inicio del router

	router.GET("/ping", func(ctx *gin.Context) {ctx.String(200, "pong")}) //nuestro router de ejemplo

	//creamos un grupo de enlaces que se dirigira a nuestros controladores
	prods := router.Group("/productos")
	{
		prods.GET("/", productHandler.GetProductos())
		prods.GET(":id", productHandler.GetProductoById())
		prods.GET("/search", productHandler.PreciosMayores())
		prods.POST("/", productHandler.Create())
	}

	router.Run(":8080")
}


//funcion que va cargar nuestro archivo .json
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