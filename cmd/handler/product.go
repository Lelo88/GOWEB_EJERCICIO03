//por aca pasamos las funcionalidades del service para tener mas controlado los datos que se pasan por request

package handler

import (
	"errors"
	"strconv"
	"strings"

	"github.com/Lelo88/GOWEB_EJERCICIO03/internal/domain"
	"github.com/Lelo88/GOWEB_EJERCICIO03/internal/product"
	"github.com/gin-gonic/gin"
)

type productHandler struct{
	product product.Service
}

func NewProductHandler(prod product.Service) *productHandler {
	return &productHandler{
		product: prod,
    }
}

func (prod *productHandler) GetProductos()gin.HandlerFunc{
	return func(ctx *gin.Context) {
		productos, err := prod.product.GetProductos()
        if err!= nil {
            ctx.JSON(500, err)
            return
        }
        ctx.JSON(200, productos)
    }
}

func (prod *productHandler) GetProductoById()gin.HandlerFunc{
	return func(ctx *gin.Context) {
        idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
        if err!= nil {
            ctx.JSON(400, gin.H{"msg": "producto invalido"})
            return
        }

		producto, err2 := prod.product.GetProductoById(id)
		if err2!= nil {
			ctx.JSON(404, gin.H{"msg": "producto no encontrado"})
			return
		}

		ctx.JSON(200, producto)
	}
}

func(prod productHandler) PreciosMayores()gin.HandlerFunc {
	return func(ctx *gin.Context) {
	precioParam := ctx.Query("price")
		price, err := strconv.ParseFloat(precioParam, 64)
		if err!= nil {
			ctx.JSON(400, gin.H{"msg": "precio invalido"})
            return
        }

		productos, err := prod.product.PreciosMayores(price)
		if err!= nil {
			ctx.JSON(404, gin.H{"msg": "precio no encontrado"})
            return
        }
		ctx.JSON(200, productos)
    }
}

func validarProducto(prod *domain.Producto) (bool, error){
	switch{
	case (prod.Name == "") || (prod.Expiration =="") || (prod.Code_Value == ""):
		return false, errors.New("los campos no pueden estar vacíos")
	case (prod.Price<=0) || (prod.Quantity <=0):
		if prod.Price<=0{
			return false, errors.New("el precio no puede ser menor o igual a cero")
		}
		if prod.Price<=0{
			return false, errors.New("la cantidad no puede ser menor o igual a cero")
		}
    }
	return true, nil
}

func validaFecha(prod *domain.Producto)(bool, error){
	
	fecha := strings.Split(prod.Expiration, "/")
	listaFechas := []int{}
	if len(fecha)!=3{
		return false, errors.New("fecha de vencimiento invalida, debe ingresar este formato: dd/mm/yyyy")
	}

	for values := range fecha{
		numero, err := strconv.Atoi(fecha[values])
		if err != nil {
			return false, errors.New("fecha de vencimiento invalida, debe ingresar numeros")
		}
		listaFechas = append(listaFechas, numero)
	}

	if (listaFechas[0]<1 || listaFechas[0]>31) || (listaFechas[1]<1 || listaFechas[1]>12) || (listaFechas[2]<2023) {
		return false, errors.New("fecha de vencimiento invalido, la fecha de vencimiento debe ser del 1/1/2023 en adelante")
	}

	return true, nil
}

func (prod *productHandler) Create() gin.HandlerFunc{
		return func(ctx *gin.Context) {
		
		var producto domain.Producto
			
		err := ctx.ShouldBindJSON(&producto)
        if err!= nil {
            ctx.JSON(400, gin.H{"error": "producto inválido"})
			return
		}

		valido, err := validarProducto(&producto)
		if !valido {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return 
		}

		valido, err = validaFecha(&producto)
		if !valido{
			ctx.JSON(400, gin.H{"error": err.Error()})
			return 
		}

		pro, err := prod.product.CreateProduct(producto)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, pro)
	}
}
/*
import (
	"strconv"

	"github.com/Lelo88/GOWEB_EJERCICIO03/internal/domain"
	"github.com/Lelo88/GOWEB_EJERCICIO03/internal/product"
	"github.com/gin-gonic/gin"
)

//funcion que me va a devolver los productos de nuestra base de datos
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
		for _, product := range domain.Producto {
			if product.Id == id {
				c.JSON(200, product)
				return
			}
		}
		c.JSON(404, gin.H{"error": "product not found"})
	}
}	

//funcion que me va a devolver los productos mayores a un precio determinado pasados por query
func PreciosMayores()gin.HandlerFunc{
	return func(c *gin.Context){
		var resultado []domain.Producto

		precio, err := strconv.ParseFloat(c.Query("price"), 64)
		
		if err!= nil {
		c.JSON(404, gin.H{"error": err.Error()})
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
		var producto domain.Producto
		
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
}*/