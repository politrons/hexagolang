package infra

import . "domain"

type OrderDAO interface {
	GetOrder(id Id) Order

	CreateOrder(Order)
}

type OrderDAOImpl struct {
	orders map[Id]Order
}

func (orderDAO OrderDAOImpl) GetOrder(id Id) Order {
	order, exist := orderDAO.orders[id]
	if !exist {
		return Order{}
	} else {
		return order
	}
}

func (orderDAO OrderDAOImpl) CreateOrder(order Order) {
	orderDAO.orders[order.Id] = order
}
