package order

import "fmt"

type MockOrderGatewayError struct {
	IOrderGateway
}

func (o MockOrderGatewayError) Save(orderData OrderDataRequest) (string, error) {
	return "", fmt.Errorf("erro ao salvar o produto")
}
