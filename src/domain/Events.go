package domain

/**
This file contains the interface and all events structs that implement that interface.

In order to be consider an event the [Process] and [Exist] function must be implemented for each
struct type.

For the [Process] function each event it will implement the logic of how to modify the Order type, to be
Rehydrated to the last state after pass all event process over him, applying Event sourcing pattern.
*/
type Event interface {
	Process(order Order) Order
	Exist(transactionId string) bool
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

/**
Just return the order previously created
*/
func (event OrderCreated) Process(order Order) Order {
	return event.Order
}

func (event OrderCreated) Exist(transactionId string) bool {
	return event.Order.Id.Value == transactionId
}

/**
Add the product in the products list inside the order
*/
func (event ProductAdded) Process(order Order) Order {
	order.Products = append(order.Products, event.Product)
	return order
}

func (event ProductAdded) Exist(transactionId string) bool {
	return event.Product.TransactionId == TransactionId{transactionId}
}

/**
filter the list of products, and create a new one without the element we want to be removed.
*/
func (event ProductRemoved) Process(order Order) Order {
	var products []Product
	for _, product := range order.Products {
		if !isSameProductAndTransaction(product, event) {
			products = append(products, product)
		}
	}
	order.Products = products
	return order
}

func (event ProductRemoved) Exist(transactionId string) bool {
	return event.Product.TransactionId == TransactionId{transactionId}

}
func isSameProductAndTransaction(product Product, event ProductRemoved) bool {
	return product.Id == event.Product.Id
}
