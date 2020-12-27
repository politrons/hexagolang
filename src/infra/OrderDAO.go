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

/**
We have the list of events attach in a map to an OrderId, in case the order was not yet created we have a nil map,
so we create one, and then we have to also check if the events already exist for that entry and if it does not
we also create one.
*/
func (orderDAO OrderDAOImpl) AddEvent(orderId OrderId, event Event) {
	if orderDAO.orderEvents == nil {
		orderDAO.orderEvents = make(map[OrderId][]Event)
	}
	events, exist := orderDAO.orderEvents[orderId]
	if !exist {
		orderDAO.orderEvents[orderId] = []Event{event}
	} else {
		orderDAO.orderEvents[orderId] = append(events, event)
	}
}
