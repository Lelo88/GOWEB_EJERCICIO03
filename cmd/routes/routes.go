package routes

import (
	"github.com/Lelo88/GOWEB_EJERCICIO03/cmd/handler"
	"github.com/Lelo88/GOWEB_EJERCICIO03/internal/domain"
	"github.com/Lelo88/GOWEB_EJERCICIO03/internal/product"
	"github.com/gin-gonic/gin"
)

type Router struct {
	db *[]domain.Producto
	en *gin.Engine //enrutador de endpoints 
}

func NewRouter(en *gin.Engine, db *[]domain.Producto) *Router {    //constructor de enlaces de endpoints
	return &Router{en: en, db: db}
}

func (r *Router) SetRoutes(){ //seteador de endpoints
	r.SetProducts()
}

func (r *Router) SetProducts(){ //seteador de controladores 
	//desde este metodo inicializamos el repositorio y servicio, es una abstraccion 
	rep:= product.NuevoRepositorio(*r.db)
	ser:= product.NewService(rep)
	prohand := handler.NewProductHandler(ser)
	productos:=r.en.Group("/productos")

	productos.GET("/ping", func(ctx *gin.Context) {ctx.String(200, "pong")})
	productos.GET("", prohand.GetProductos())
	productos.GET("/:id", prohand.GetProductoById())
	productos.GET("/search", prohand.PreciosMayores())
	productos.POST("", prohand.Create())
}