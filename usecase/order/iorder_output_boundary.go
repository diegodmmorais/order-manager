package order

type IOrderOutputBoundary interface {
	Success(orderInput OrderInputData) (*OrderResponse, error)
}
