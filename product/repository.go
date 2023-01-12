package product

type Producto struct {
	Id 				int		`json:"id"`
	Name 			string 	`json:"name" binding:"required"`
	Quantity 		int 	`json:"quantity" binding:"required"`
	Code_Value 		string 	`json:"code_value" binding:"required"`
	Is_Published 	bool 	`json:"is_published"`
	Expiration 		string 	`json:"expiration" binding:"required"`
	Price 			float64 `json:"price" binding:"required"`
}