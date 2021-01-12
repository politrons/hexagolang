package domain

type TransactionId struct{ Value string }

type ProductId struct{ Value string }

type OrderId struct{ Value string }

type Name struct{ Value string }

type Description struct{ Value string }

type Price struct{ Value float64 }

type Order struct {
	Id         OrderId
	Products   []Product
	TotalPrice Price
}

type ProductI interface {
	HasProductSameTransactionId(transactionId string) bool
}

type Product struct {
	TransactionId TransactionId
	Id            ProductId
	Name          Name
	Price         Price
	Description   Description
}

func (product Product) HasProductSameTransactionId(transactionId string) bool {
	return product.TransactionId == TransactionId{transactionId}
}
