package order

type MockOrderGateway struct {
	IOrderGateway
}

func (o MockOrderGateway) Save(orderData OrderDataRequest) (string, error) {
	return "#1", nil
}