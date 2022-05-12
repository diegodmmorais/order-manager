package product

import "github.com/stretchr/testify/mock"

type MockProductGateway struct {
	mock.Mock
}

func (mock *MockProductGateway) FindByProduct(productID string) (ProductOutputData, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(ProductOutputData), args.Error(1)
}
