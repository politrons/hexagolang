package handler

import (
	"app/command"
	. "domain"
	"infra/dao"
)

type OrderHandler interface {
	CreateOrder(orderId string, command command.CreateOrderCommand)
}

type OrderHandlerImpl struct {
	OrderDAO dao.OrderDAO
}

func (handler OrderHandlerImpl) CreateOrder(orderId string, command command.CreateOrderCommand) {
	exist := handler.eventAlreadyExist(command, orderId)
	if !exist {
		order := Order{
			Id:         OrderId{Value: command.Id},
			Products:   []Product{},
			TotalPrice: Price{},
		}
		orderCreated := OrderCreated{Order: order}
		handler.OrderDAO.AddEvent(order.Id, orderCreated)
	}
}

func (handler OrderHandlerImpl) eventAlreadyExist(
	orderCommand command.CreateOrderCommand,
	orderId string) bool {
	var exist = false
	for _, event := range handler.OrderDAO.GetEvents(OrderId{Value: orderCommand.Id}) {
		if event.Exist(orderId) {
			exist = true
		}
	}
	return exist
}
