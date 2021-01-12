package domain

/**
This file contains the interface and all events structs that implement that interface.

In order to be consider an [Event] the [Process] and [Exist] function must be implemented for each
struct type.

For the [Process] function each event it will implement the logic of how to modify the Order type, to be
Rehydrated to the last state after pass all event process over him, applying [Event Sourcing] pattern.
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

/**
Check if the transaction used form the operation is the same that was used to create the Id of the Order
*/
func (event OrderCreated) Exist(transactionId string) bool {
	return event.Order.Id.Value == transactionId
}

func (event OrderCreated) GetProduct() (bool, Product) {
	return false, Product{}
}

/**
Add the product in the products list inside the order

Also we increase the price of the product in the [TotalPrice] of the [Order]
*/
func (event ProductAdded) Process(order Order) Order {
	order.Products = append(order.Products, event.Product)
	order.TotalPrice = Price{order.TotalPrice.Value + event.Product.Price.Value}
	return order
}

func (event ProductAdded) GetProduct() (bool, Product) {
	return true, event.Product
}

/**
Check if the transaction used form the operation is the same that was used to create the TransactionId of the Product
*/
func (event ProductAdded) Exist(transactionId string) bool {
	return event.Product.HasProductSameTransactionId(transactionId)
}

/**
Filter the list of products, and create a new one without the element we want to be removed.

Also we subtract the price of the product in the [TotalPrice] of the [Order]
*/
func (event ProductRemoved) Process(order Order) Order {
	var products []Product
	for _, product := range order.Products {
		if !isSameProductId(product, event) {
			products = append(products, product)
		} else {
			order.TotalPrice = Price{order.TotalPrice.Value - product.Price.Value}
		}
	}
	order.Products = products
	return order
}

func (event ProductRemoved) Exist(transactionId string) bool {
	return event.Product.HasProductSameTransactionId(transactionId)
}

func (event ProductRemoved) GetProduct() (bool, Product) {
	return true, event.Product
}

func isSameProductId(product Product, event ProductRemoved) bool {
	return product.Id == event.Product.Id
}
