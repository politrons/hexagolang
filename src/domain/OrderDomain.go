package domain

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

type Product struct {
	Id          ProductId
	Name        Name
	Price       Price
	Description Description
}
