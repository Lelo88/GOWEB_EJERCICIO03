package main

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//creo la estructura para traerme los datos del archivo json
type Producto struct {
	Id 				int		`json:"id"`
	Name 			string 	`json:"name" binding:"required"`
	Quantity 		int 	`json:"quantity" binding:"required"`
	Code_Value 		string 	`json:"code_value" binding:"required"`
	Is_Published 	bool 	`json:"is_published"`
	Expiration 		string 	`json:"expiration" binding:"required"`
	Price 			float64 `json:"price" binding:"required"`
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

//funcion que me va a devolver un producto por ID. Si no existe, return 404
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

//funcion que me va a devolver un listado con los productos con precio mayor al colocado 
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

//--------FUNCIONES PARA EL METODO POST-------

//funcion que va a validar un objeto json vacio
func ValidarProducto(producto *Producto) (bool,error){
	switch {
	case producto.Name == "" || producto.Code_Value == "" || producto.Expiration == "":
		return false, errors.New("los campos no pueden estar vacios")
	case producto.Price<=0:
		return false, errors.New("el precio no puede ser menor o igual a 0")
	case producto.Quantity<=0:
		return false, errors.New("la cantidad no puede ser menor o igual a 0")
    }
	return true, nil
}

//funcion de validacion de fecha 
func ValidarFecha(product *Producto) (bool,error){

	fecha := strings.Split(product.Expiration, "/")
	listaFechas := []int{}
	if len(fecha)!=3{
		return false, errors.New("fecha de vencimiento invalida, debe ingresar este formato: dd/mm/yyyy")
	}

	for values := range fecha{
		numero, err := strconv.Atoi(fecha[values])
		if err != nil {
			return false, errors.New("invalid expiration date, must be numbers")
		}
		listaFechas = append(listaFechas, numero)
	}

	if (listaFechas[0]<1 || listaFechas[0]>31) && (listaFechas[1]<1 || listaFechas[1]>12) && (listaFechas[2]<2023){
		return false, errors.New("fecha de vencimiento invalido, la fecha de vencimiento debe ser del 1/1/2023 en adelante")
	}

	return true, nil
}

//funcion que va a validar el codigo de id

func ValidaCodigoID(codigo string) bool{
	for _,producto := range productos {
		if producto.Code_Value == codigo{
            return false
        }
	}
	return true
}
//2. funcion que va a ingresar un producto ingresado por request

func CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
        
		//si lo que ingresamos no es un objeto json, nos devuelve un error
		var product Producto
		
		err := c.ShouldBindJSON(&product)
        if err!= nil {
            c.JSON(400, gin.H{"error": "producto inválido"})
			return
		}

		//si lo que ingresamos no es valido, nos devuelve un error
		valido, err := ValidarProducto(&product)
		if !valido {
			c.JSON(400, gin.H{"error": err.Error()})
			return 
		}
		
		valido, err = ValidarFecha(&product)
		if !valido{
			c.JSON(400, gin.H{"error": err.Error()})
			return 
		}
		
		valido = ValidaCodigoID(product.Code_Value)
		if !valido {
			c.JSON(400, gin.H{"error": "codigo existente"})
            return
        }

		product.Id = len(productos) + 1
		productos = append(productos, product)
		c.JSON(201, product)
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
		products.GET("", getProductos()) //localhost:8080/products --> si colocamos mal algo, nos va a tirar el error indicado
		products.GET("/:id", getProductoById()) //localhost:8080/products/numero entero --> si colocamos mal algo, nos va a tirar el error indicado
		products.GET("/search", preciosMayores()) //localhost:8080/products/search?price=900 --> si colocamos mal algo, nos va a tirar el error indicado
		products.POST("", CreateProduct()) //2. llamamos al metodo post para crear un producto, en caso de algun dato erroneo, evalua el error
		//añadimos el  metodo POST para la segun parte del ejercicio

	}	

    router.Run()
}