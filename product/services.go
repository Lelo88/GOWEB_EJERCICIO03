package product

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

var Productos = []Producto{}

func LoadProducts(path string) {
	file, err := os.ReadFile(path)
    if err!= nil {
        panic(err)
    }

	err = json.Unmarshal(file, &Productos)
    if err!= nil {
		panic(err)
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

func ValidaCodigoID(codigo string) bool{
	for _,producto := range Productos {
		if producto.Code_Value == codigo{
            return false
        }
	}
	return true
}