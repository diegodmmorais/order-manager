package gateway

import "github.com/stretchr/testify/mock"

type MockOrderRepository struct {
	mock.Mock
}

func (mock *MockOrderRepository) Save(orderMapper OrderInputMapper) (string, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(string), args.Error(1)
}
