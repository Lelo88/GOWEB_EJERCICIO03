// estructura que representa a nuestra base de datos extraida con los datos del archivo products.json
package product

import (
	"errors"

	"github.com/Lelo88/GOWEB_EJERCICIO03/internal/domain"
)

type Repository interface {
    GetProductos() ([]domain.Producto, error)
	GetProductoByID(id int) (domain.Producto, error)
	PreciosMayores(precio float64) []domain.Producto
	CreateProduct(p domain.Producto) (domain.Producto, error)
}

//estructura que simula ser un repositorio
type repositorio struct {
	productos []domain.Producto
}

func NuevoRepositorio(productos []domain.Producto) *repositorio {
	return &repositorio{
        productos: productos,
    }
}

func (rep *repositorio) GetProductos() []domain.Producto {
	return rep.productos
}

func (rep *repositorio) GetProductoByID(id int) (domain.Producto, error){
	for _, producto := range rep.productos {
        if producto.Id == id {
            return producto, nil
        }
    }
    return domain.Producto{}, errors.New("no se encuentra el producto solicitado")
}

func (rep *repositorio) PreciosMayores(precio float64) ([]domain.Producto){

	var resultado []domain.Producto
	for _, producto := range rep.productos {
		if producto.Price == precio {
            resultado = append(resultado, producto)
        }
    }
	return resultado
}

func (rep *repositorio) CreateProduct(p domain.Producto) (domain.Producto, error) {
	if !rep.ValidaProducto(p.Code_Value) {
		return domain.Producto{}, errors.New("producto existente")
    }
	
	p.Id = len(rep.productos) + 1
	rep.productos = append(rep.productos, p)
	return p, nil
}

func (rep *repositorio) ValidaProducto(codigo string) bool{
	for _, producto :=range rep.productos{
		if producto.Code_Value == codigo {
			return false
		}
	}
	return true
}