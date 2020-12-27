package response

import . "domain"

type OrderResponse struct {
	Exist bool
	Order Order
}
