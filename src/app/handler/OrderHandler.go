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

func (handler OrderHandlerImpl) CreateOrder(transactionId string, command command.CreateOrderCommand) {
	exist := handler.eventAlreadyExist(command, transactionId)
	if !exist {
		order := Order{
			Id:         OrderId{Value: command.Id},
			Products:   []Product{},
			TotalPrice: Price{},
		}
		orderCreated := OrderCreated{TransactionId: transactionId, Order: order}
		handler.OrderDAO.AddEvent(order.Id, orderCreated)
	}
}

func (handler OrderHandlerImpl) eventAlreadyExist(
	orderCommand command.CreateOrderCommand,
	transactionId string) bool {
	var exist = false
	for _, event := range handler.OrderDAO.GetEvents(OrderId{Value: orderCommand.Id}) {
		if event.Exist(transactionId) {
			exist = true
		}
	}
	return exist
}
