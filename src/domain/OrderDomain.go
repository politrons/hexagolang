package domain

type Id struct{ Value int }

type Name struct{ Value string }

type Description struct{ Value string }

type Price struct{ Value float64 }

type Order struct {
	Id          Id
	Name        Name
	Price       Price
	Description Description
}
