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

func (handler ProductHandlerImpl) AddProduct(command command.AddProductCommand) {
	product := Product{
		Id:          ProductId{Value: command.Id},
		Price:       Price{Value: command.Price},
		Description: Description{Value: command.Description},
	}

	productAdded := ProductAdded{Product: product}
	handler.OrderDAO.AddEvent(OrderId{Value: command.OrderId}, productAdded)
}
