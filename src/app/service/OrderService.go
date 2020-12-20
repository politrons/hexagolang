package service

import (
	. "app/command"
	. "domain"
	"infra"
)

type OrderService interface {
	GetOrder(id int) Order
}

type OrderServiceImpl struct {
	OrderDAO infra.OrderDAO
}

func (service OrderServiceImpl) GetOrder(id int) Order {
	return service.OrderDAO.GetOrder(Id{id})
}

func (service OrderServiceImpl) CreateOrder(command CreateOrderCommand) Order {
	order := Order{
		Id:          Id{Value: command.Id},
		Name:        Name{Value: command.Name},
		Price:       Price{Value: command.Price},
		Description: Description{Value: command.Description},
	}
	return order
}
