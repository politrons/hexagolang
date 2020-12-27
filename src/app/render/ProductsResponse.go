package render

import . "domain"

type ProductsResponse struct {
	TransactionId string
	Products      []Product
}
