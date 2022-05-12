package order

type IOrderPresenter interface {
	Success(orderInput OrderSuccessInputData) (*OrderResponse, error)
}
