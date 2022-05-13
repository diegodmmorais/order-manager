package usecase

type IOrderGateway interface {
	Save(orderData OrderInputData) (string, error)
}
