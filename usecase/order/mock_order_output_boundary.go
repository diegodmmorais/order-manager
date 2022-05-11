package order

type MockOrderOutputBoundary struct {
	IOrderOutputBoundary
}

func (m MockOrderOutputBoundary) Success(orderInput OrderInputData) *OrderResponse {
	return &OrderResponse{
		ID:           orderInput.OrderID,
		Message:      "Pedido Salvo",
		CustomerName: orderInput.CustomerName,
	}
}
