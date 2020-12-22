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
	orderDAO infra.OrderDAO
}

func (handler OrderHandlerImpl) CreateOrder(command command.CreateOrderCommand) {
	order := Order{
		Id:          Id{Value: command.Id},
		Name:        Name{Value: command.Name},
		Price:       Price{Value: command.Price},
		Description: Description{Value: command.Description},
	}
	handler.orderDAO.CreateOrder(order)
}
