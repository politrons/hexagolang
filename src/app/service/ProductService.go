package service

import (
	. "domain"
	"infra/dao"
)

/**
Service layer [interface] where we define the API of this Service.
In order to have an implementation of this interface you need to have a [struct] which
you extend methods like the one defines in the interface
*/
type ProductService interface {
	GetProduct(id string) (bool, Product)

	GetAllProduct() []Product
}

/**
Implementation type of interface [OrderService].
To be consider a interface implementation you need also to create extended functions of this type,
that implement the interface methods.
*/
type ProductServiceImpl struct {
	ProductDAO dao.ProductDAO
}

func (service ProductServiceImpl) GetProduct(productId string) (bool, Product) {
	return service.ProductDAO.FindProduct(productId)
}

func (service ProductServiceImpl) GetAllProduct() []Product {
	return service.ProductDAO.GetAllProducts()
}
