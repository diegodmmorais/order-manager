package order

type IOrderGateway interface {
	Save(orderData OrderInputData) (string, error)
}
