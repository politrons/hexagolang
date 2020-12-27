package infra

import . "domain"

type ProductDAO interface {
	FindProduct(productId string) (bool, Product)

	GetAllProducts() []Product
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

/**
Function to return some mock products
*/
func (productDAO ProductDAOImpl) GetAllProducts() []Product {
	return []Product{
		{Id: ProductId{Value: "1981"},
			Name:        Name{Value: "Coke-cole"},
			Price:       Price{Value: 2.50},
			Description: Description{Value: "Sugar soda"},
		},
		{Id: ProductId{Value: "1982"},
			Name:        Name{Value: "Twix"},
			Price:       Price{Value: 1.50},
			Description: Description{Value: "Chocolate Candy bar"},
		},
	}
}
