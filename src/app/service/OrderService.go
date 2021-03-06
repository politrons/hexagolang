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
type OrderService interface {
	GetOrder(id string) chan Order
}

/**
Implementation type of interface [OrderService].
To be consider a interface implementation you need also to create extended functions of this type,
that implement the interface methods.
*/
type OrderServiceImpl struct {
	OrderDAO dao.OrderDAO
}

func (service OrderServiceImpl) GetOrder(id string) chan Order {
	return service.OrderDAO.Rehydrate(OrderId{Value: id})

}
