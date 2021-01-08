package handler

import (
	"app/command"
	. "domain"
	"infra/dao"
)

type ProductHandler interface {
	AddProduct(transactionId string, command command.AddProductCommand)

	RemoveProduct(transactionId string)
}

type ProductHandlerImpl struct {
	OrderDAO dao.OrderDAO
}

/**
Handler method that transform the Command to add product into ProductAdded event.
This event contains all the information to change the state of the order
*/
func (handler ProductHandlerImpl) AddProduct(transactionId string, command command.AddProductCommand) {
	exist := handler.eventAlreadyExist(OrderId{Value: command.OrderId}, transactionId)
	if !exist {
		product := Product{
			TransactionId: TransactionId{Value: transactionId},
			Id:            ProductId{Value: command.Id},
			Price:         Price{Value: command.Price},
			Description:   Description{Value: command.Description},
		}
		productAdded := ProductAdded{Product: product}
		handler.OrderDAO.AddEvent(OrderId{Value: command.OrderId}, productAdded)
		handler.OrderDAO.Rehydrate(OrderId{Value: command.OrderId})
	}
}

func (handler ProductHandlerImpl) RemoveProduct(transactionId string, command command.RemoveProductCommand) {
	exist := handler.eventAlreadyExist(OrderId{Value: command.OrderId}, transactionId)
	if !exist {
		product := Product{
			TransactionId: TransactionId{Value: transactionId},
			Id:            ProductId{Value: command.Id},
			Price:         Price{},
			Description:   Description{},
		}
		productRemoved := ProductRemoved{Product: product, TransactionId: TransactionId{Value: transactionId}}
		handler.OrderDAO.AddEvent(OrderId{Value: command.OrderId}, productRemoved)
		handler.OrderDAO.Rehydrate(OrderId{Value: command.OrderId})
	}
}

func (handler ProductHandlerImpl) eventAlreadyExist(orderId OrderId,
	transactionId string) bool {
	var exist = false
	for _, event := range handler.OrderDAO.GetEvents(orderId) {
		if event.Exist(transactionId) {
			exist = true
		}
	}
	return exist
}
