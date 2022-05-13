package gateway

import "github.com/stretchr/testify/mock"

type MockCustomerApi struct {
	mock.Mock
}

func (mock *MockCustomerApi) FindByCustomer(customerID string) (CustomerOutPutMapper, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(CustomerOutPutMapper), args.Error(1)
}
