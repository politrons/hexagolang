package handler

import (
	"app/command"
	. "domain"
	"infra/dao"
)

type OrderHandler interface {
	CreateOrder(command command.CreateOrderCommand)
}

type OrderHandlerImpl struct {
	OrderDAO dao.OrderDAO
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
