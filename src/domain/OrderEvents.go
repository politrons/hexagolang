package domain

type Event interface {
	process(order Order) interface{}
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

func (event OrderCreated) process(order Order) interface{} {
	return event.Order
}

func (event ProductAdded) process(order Order) interface{} {
	products := append(order.products, event.Product)
	return products
}

func (event ProductRemoved) process(order Order) interface{} {
	var products []Product
	for _, product := range order.products {
		if product.Id != event.Product.Id {
			products = append(products, product)
		}
	}
	return products
}
