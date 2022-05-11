package order

type MockOrderOutputBoundary struct {
	IOrderOutputBoundary
}

func (m MockOrderOutputBoundary) Success(orderInput OrderInputData) *OrderResponse {
	return &OrderResponse{
		ID:           orderInput.OrderID,
		Message:      "pedido salvo com sucesso",
		CustomerName: orderInput.CustomerName,
	}
}
