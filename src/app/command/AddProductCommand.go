package command

type AddProductCommand struct {
	OrderId     string
	Id          string
	Price       float64
	Description string
}
