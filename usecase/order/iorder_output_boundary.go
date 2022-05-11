package order

type IOrderOutputBoundary interface {
	success(order OrderInputData) *OrderResponse
}
