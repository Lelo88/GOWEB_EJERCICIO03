package handler

import (
	"strconv"

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

		producto, err := prod.product.GetProductoById(id)
		if err!= nil {
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