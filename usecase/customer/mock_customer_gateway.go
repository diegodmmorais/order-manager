package customer

import "github.com/stretchr/testify/mock"

type MockCustomerGateway struct {
	mock.Mock
}

func (mock MockCustomerGateway) FindByCustomer(customerID string) (CustomerOutputData, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(CustomerOutputData), args.Error(1)
}
