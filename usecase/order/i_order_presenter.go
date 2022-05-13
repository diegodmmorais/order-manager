package usecase

type IOrderPresenter interface {
	Success(orderInput OrderSuccessInputData) (*OrderResponse, error)
}
