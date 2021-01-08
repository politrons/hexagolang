package domain

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
	TransactionId TransactionId
	Product       Product
}

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
	return product.Id == event.Product.Id && event.TransactionId.Value == product.TransactionId.Value
}
