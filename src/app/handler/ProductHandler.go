package handler

import (
	"app/command"
	. "domain"
	"infra"
)

type ProductHandler interface {
	AddProduct(command command.AddProductCommand)
}

type ProductHandlerImpl struct {
	OrderDAO infra.OrderDAO
}

/**
Handler method that transform the Command to add product into ProductAdded event.
This event contains all the information to change the state of the order
*/
func (handler ProductHandlerImpl) AddProduct(command command.AddProductCommand) {
	product := Product{
		Id:          ProductId{Value: command.Id},
		Price:       Price{Value: command.Price},
		Description: Description{Value: command.Description},
	}

	productAdded := ProductAdded{Product: product}
	handler.OrderDAO.AddEvent(OrderId{Value: command.OrderId}, productAdded)
}
