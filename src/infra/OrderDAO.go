package infra

import . "domain"

type OrderDAO interface {
	Rehydrate(orderId OrderId) Order

	AddEvent(orderId OrderId, event Event)

	/*	GetOrder(id OrderId) Order

		CreateOrder(orderCreated OrderCreated)*/
}

type OrderDAOImpl struct {
	orderEvents map[OrderId][]Event
}

func (orderDAO OrderDAOImpl) AddEvent(orderId OrderId, event Event) {
	events := orderDAO.orderEvents[orderId]
	orderDAO.orderEvents[orderId] = append(events, event)
}

func (orderDAO OrderDAOImpl) Rehydrate(orderId OrderId) Order {

}

/*
func (orderDAO OrderDAOImpl) GetOrder(id OrderId) Order {
	order, exist := orderDAO.orders[id]
	if !exist {
		return Order{}
	} else {
		return order
	}
}

func (orderDAO OrderDAOImpl) CreateOrder(orderCreated OrderCreated) {
	orderDAO.orders[orderCreated.Order.Id] = orderCreated.Order
}*/
