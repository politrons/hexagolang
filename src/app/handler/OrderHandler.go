package handler

import (
	"app/command"
	"infra"
)

type OrderHandler interface {
	CreateOrder(command command.CreateOrderCommand)
}

type OrderHandlerImpl struct {
	orderDAO infra.OrderDAO
}

func (handler OrderHandlerImpl) CreateOrder(command command.CreateOrderCommand) {
	handler.orderDAO.CreateOrder()
}
