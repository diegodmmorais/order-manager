package gateway

import "github.com/stretchr/testify/mock"

type MockProductApi struct {
	mock.Mock
}

func (mock *MockProductApi) FindByProduct(productID string) (ProductOutputMapper, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(ProductOutputMapper), args.Error(1)
}
