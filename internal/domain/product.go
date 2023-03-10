//reformulo la estructura a partir de aca utilizando metodos para la bbdd y para los servicios solicitados

package domain

type Producto struct {
	Id 				int		`json:"id"`
	Name 			string 	`json:"name" validate:"required"`
	Quantity 		int 	`json:"quantity" validate:"required"`
	Code_Value 		string 	`json:"code_value" validate:"required"`
	Is_Published 	bool 	`json:"is_published"`
	Expiration 		string 	`json:"expiration" validate:"required"`
	Price 			float64 `json:"price" validate:"required"`
}