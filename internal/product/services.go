package product

import (

	"github.com/Lelo88/GOWEB_EJERCICIO03/internal/domain"
)

type Service interface {
	GetProductos() ([]domain.Producto, error)
	GetProductoById(id int) (domain.Producto, error)
	PreciosMayores(precio float64) []domain.Producto
	CreateProduct(p domain.Producto) (domain.Producto, error)
}

type servicio struct {
	rep Repository
}

func NewService(re Repository) Service {
    return &servicio{re}
}

func (s *servicio) GetProductos() ([]domain.Producto, error) {
	productos, err:= s.rep.GetProductos()
	if err!= nil {
		return nil, err
	}
	return productos, nil
}

func (s *servicio) GetProductoById(id int) (domain.Producto, error){
	producto, err := s.rep.GetProductoByID(id)
	if err!= nil {
		return domain.Producto{}, err
	}
	return producto, nil
}

func (s *servicio) PreciosMayores(precio float64) []domain.Producto{
	productos := s.rep.PreciosMayores(precio)
	if len(productos) == 0 {
		return []domain.Producto{}
    }
	return productos
}

func (s *servicio) CreateProduct(p domain.Producto) (domain.Producto, error){
	producto, err := s.rep.CreateProduct(p)
	if err!= nil {
        return domain.Producto{}, err
    }
	return producto, nil
}


//funcion que nos carga el archivo json a la estructura producto

/*

//--------FUNCIONES PARA EL METODO POST-------

//funcion que va a validar un objeto json vacio
func ValidarProducto(producto *Producto) (bool,error) { 
	switch {
	case (producto.Name == "") || (producto.Code_Value == "") || (producto.Expiration == ""):
		return false, errors.New("los campos no pueden estar vacios")
	case producto.Price<=0 || producto.Quantity<=0:
		if producto.Price <=0 {
			return false, errors.New("el precio no puede ser menor o igual a 0")
		}
		if producto.Quantity <=0 {
			return false, errors.New("la cantidad no puede ser menor o igual a 0")
		}
	}
	return true, nil
}

//funcion que valida la fecha de vencimiento de un producto
func ValidarFecha(product *Producto) (bool,error){

	fecha := strings.Split(product.Expiration, "/")
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

//funcion que valida nuestro codigo de mercaderia
func ValidaCodigoID(codigo string) bool{
	for _,producto := range Productos {
		if producto.Code_Value == codigo{
            return false
        }
	}
	return true
}*/