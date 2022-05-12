package order

import (
	"github.com/stretchr/testify/mock"
)

type MockOrderGateway struct {
	mock.Mock
}

func (mock MockOrderGateway) Save(orderData OrderInputData) (string, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(string), args.Error(1)
}
