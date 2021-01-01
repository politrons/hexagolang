package handler

import (
	"app/command"
	. "domain"
	"infra/dao"
)

type ProductHandler interface {
	AddProduct(transactionId string, command command.AddProductCommand)
}

type ProductHandlerImpl struct {
	OrderDAO dao.OrderDAO
}

/**
Handler method that transform the Command to add product into ProductAdded event.
This event contains all the information to change the state of the order
*/
func (handler ProductHandlerImpl) AddProduct(transactionId string, command command.AddProductCommand) {
	exist := handler.eventAlreadyExist(command, transactionId)
	if !exist {
		product := Product{
			TransactionId: TransactionId{Value: transactionId},
			Id:            ProductId{Value: command.Id},
			Price:         Price{Value: command.Price},
			Description:   Description{Value: command.Description},
		}
		productAdded := ProductAdded{Product: product}
		handler.OrderDAO.AddEvent(OrderId{Value: command.OrderId}, productAdded)
	}
}

func (handler ProductHandlerImpl) eventAlreadyExist(
	addProductCommand command.AddProductCommand,
	transactionId string) bool {
	var exist = false
	for _, event := range handler.OrderDAO.GetEvents(OrderId{Value: addProductCommand.Id}) {
		if event.Exist(transactionId) {
			exist = true
		}
	}
	return exist
}
