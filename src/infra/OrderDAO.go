package infra

import . "domain"

type OrderDAO interface {
	Rehydrate(orderId OrderId) Order

	AddEvent(orderId OrderId, event Event)
}

type OrderDAOImpl struct {
	orderEvents map[OrderId][]Event
}

func (orderDAO OrderDAOImpl) AddEvent(orderId OrderId, event Event) {
	events := orderDAO.orderEvents[orderId]
	orderDAO.orderEvents[orderId] = append(events, event)
}

func (orderDAO OrderDAOImpl) Rehydrate(orderId OrderId) Order {
	var order = Order{}
	for _, event := range orderDAO.orderEvents[orderId] {
		order = event.Process(order)
	}
	return order
}
