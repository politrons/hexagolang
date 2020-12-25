package domain

type Event interface {
	Process(order Order) Order
}

type OrderCreated struct {
	Order Order
}

type ProductAdded struct {
	Product Product
}

type ProductRemoved struct {
	Product Product
}

func (event OrderCreated) Process(order Order) Order {
	return event.Order
}

func (event ProductAdded) Process(order Order) Order {
	order.products = append(order.products, event.Product)
	return order
}

func (event ProductRemoved) Process(order Order) Order {
	var products []Product
	for _, product := range order.products {
		if product.Id != event.Product.Id {
			products = append(products, product)
		}
	}
	order.products = products
	return order
}
