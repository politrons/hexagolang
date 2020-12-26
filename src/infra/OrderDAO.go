package infra

import . "domain"

type OrderDAO interface {
	Rehydrate(orderId OrderId) (bool, Order)

	AddEvent(orderId OrderId, event Event)
}

type OrderDAOImpl struct {
	orderEvents map[OrderId][]Event
}

/**
Rehydrate function is the keystone of event sourcing. We can recreate from scratch the state of an entity
using all the events that just happens from the creation of the order, and then all the products added, or deleted
afterwards.
Each Event must implement [Process] which it will interact with Order model to recreate to the latest state.
*/
func (orderDAO OrderDAOImpl) Rehydrate(orderId OrderId) (bool, Order) {
	var exist = false
	var order = Order{}
	for _, event := range orderDAO.orderEvents[orderId] {
		order = event.Process(order)
		exist = true
	}
	return exist, order
}

func (orderDAO OrderDAOImpl) AddEvent(orderId OrderId, event Event) {
	events := orderDAO.orderEvents[orderId]
	orderDAO.orderEvents[orderId] = append(events, event)
}
