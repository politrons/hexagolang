package dao

import (
	. "domain"
	"log"
)

type OrderDAO interface {
	Create() OrderDAO

	GetEvents(orderId OrderId) []Event

	Rehydrate(orderId OrderId) chan Order

	AddEvent(orderId OrderId, event Event)
}

type OrderDAOImpl struct {
	OrderEvents map[OrderId][]Event
}

/**
Create function, it initialize the [OrderEvents] Map and return the instance of OrderDAO
*/
func (orderDAO OrderDAOImpl) Create() OrderDAO {
	orderDAO.OrderEvents = make(map[OrderId][]Event)
	return orderDAO
}

/**
GetEvents function, it return the events for a specific OrderId
*/
func (orderDAO OrderDAOImpl) GetEvents(orderId OrderId) []Event {
	return orderDAO.OrderEvents[orderId]
}

/**
Rehydrate function is the keystone of event sourcing. We can recreate from scratch the state of an entity
using all the events that just happens from the creation of the order, and then all the products added, or deleted
afterwards.
Each Event must implement [Process] which it will interact with Order model to recreate to the latest state.
*/
func (orderDAO OrderDAOImpl) Rehydrate(orderId OrderId) chan Order {
	channel := make(chan Order)
	go func() {
		var order = Order{}
		for _, event := range orderDAO.OrderEvents[orderId] {
			order = event.Process(order)
		}
		channel <- order
	}()
	return channel
}

/**
AddEvent function, we have the list of events attach in a map to an OrderId, in case the order was not yet created we have a nil map,
so we create one, and then we have to also check if the events already exist for that entry and if it does not
we also create one.
*/
func (orderDAO OrderDAOImpl) AddEvent(orderId OrderId, event Event) {
	if orderDAO.OrderEvents == nil {
		orderDAO.OrderEvents = make(map[OrderId][]Event)
	}
	events, exist := orderDAO.OrderEvents[orderId]
	if !exist {
		orderDAO.OrderEvents[orderId] = []Event{event}
		log.Printf("Creating new Order events map %s!", orderDAO.OrderEvents)
	} else {
		orderDAO.OrderEvents[orderId] = append(events, event)
		log.Printf("Adding event in current Order with id %s!", orderDAO.OrderEvents)
	}
}
