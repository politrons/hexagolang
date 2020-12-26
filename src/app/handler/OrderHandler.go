package handler

import (
	"app/command"
	. "domain"
	"infra"
)

type OrderHandler interface {
	CreateOrder(command command.CreateOrderCommand)
}

type OrderHandlerImpl struct {
	OrderDAO infra.OrderDAO
}

func (handler OrderHandlerImpl) CreateOrder(command command.CreateOrderCommand) {
	order := Order{
		Id:         OrderId{Value: command.Id},
		Products:   []Product{},
		TotalPrice: Price{},
	}

	orderCreated := OrderCreated{Order: order}
	handler.OrderDAO.AddEvent(order.Id, orderCreated)
}
