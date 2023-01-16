//siempre inicia todo desde aca yendo a services 
package product

import (
	"errors"
	"github.com/Lelo88/GOWEB_EJERCICIO03/internal/domain"
)

//se crea la interfaz de repository para crear los metodos que estaran asociados 
type Repository interface {
    GetProductos() ([]domain.Producto, error)
	GetProductoByID(id int) (domain.Producto, error)
	PreciosMayores(precio float64) []domain.Producto
	CreateProduct(p domain.Producto) (domain.Producto, error)
} 

//estructura que simula ser un repositorio
type repositorio struct {
	productos []domain.Producto
} //se crea a partir de nuestra "base de datos" inicializada en una estructura que se encuentra en la carpeta domain

func NuevoRepositorio(listproductos []domain.Producto) Repository {
	return &repositorio{listproductos}
} //creacion del nuevo repositorio a partir de un archivo correspondiente

func (rep *repositorio) GetProductos() ([]domain.Producto, error) {
	return rep.productos, nil
} //este metodo obtendra todos los productos que guardamos en memoria

func (rep *repositorio) GetProductoByID(id int) (domain.Producto, error){
	for _, producto := range rep.productos {
        if producto.Id == id {
            return producto, nil
        }
    }
    return domain.Producto{}, errors.New("no se encuentra el producto solicitado")
} //este metodo me devolvera el producto que 

func (rep *repositorio) PreciosMayores(precio float64) ([]domain.Producto){

	var resultado []domain.Producto
	for _, producto := range rep.productos {
		if producto.Price > precio {
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