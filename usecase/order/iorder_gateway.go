package order

type IOrderGateway interface {
	Save(orderData OrderDataRequest) (string, error)
}
