package Views

type OrderView struct {
	Slug string `json:"id"`
	ProductId string `json:"product_id"`
	Quantity int `json:"quantity"`
	status string `json:"status"`
}

func GetOrderView (order *Order) (orderView *OrderView) {
	orderView = &OrderView{}
	orderView.Slug = order.Slug
	orderView.ProductId = order.Slug
	orderView.Quantity = order.Quantity
	orderView.status = order.GetOrderStatusAsString()
	return orderView
}
