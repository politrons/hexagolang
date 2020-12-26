package infra

import . "domain"

type ProductDAO interface {
	FindProduct(productId string) (bool, Product)
}

type ProductDAOImpl struct {
	products map[string]Product
}

/**
We search the Product from the map of products, in case it does not exist, we receive the bool flag as false
*/
func (productDAO ProductDAOImpl) FindProduct(productId string) (bool, Product) {
	product, exist := productDAO.products[productId]
	if exist {
		return true, product
	}
	return false, product
}
