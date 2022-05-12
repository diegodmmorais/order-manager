package customer

import (
	"testing"

	"github.com/diego-dm-morais/order-manager/usecase/customer"
	"github.com/stretchr/testify/assert"
)

func Test_buscando_um_cliente(t *testing.T) {

	mockCustomerApi := new(MockCustomerApi)
	var gateway customer.ICustomerGateway = CreateCustomerGateway(mockCustomerApi)

	mockCustomerApi.On("FindByCustomer").Return(CustomerOutPutMapper{
		Name:                   "Diego Morais",
		Telephone:              "+55 015 11 9 9999-9999",
		IdentificationDocument: "000.000.000-00",
	}, nil)

	customerResponse, err := gateway.FindByCustomer("#1")

	assert.NotNil(t, customerResponse)
	assert.Nil(t, err)
	assert.NotNil(t, customerResponse.Name)
	assert.NotEmpty(t, customerResponse.Name)
	assert.Equal(t, "Diego Morais", customerResponse.Name)
	assert.NotNil(t, customerResponse.IdentificationDocument)
	assert.NotEmpty(t, customerResponse.IdentificationDocument)
	assert.Equal(t, "000.000.000-00", customerResponse.IdentificationDocument)
	assert.NotNil(t, customerResponse.Telephone)
	assert.NotEmpty(t, customerResponse.Telephone)
	assert.Equal(t, "+55 015 11 9 9999-9999", customerResponse.Telephone)
}
