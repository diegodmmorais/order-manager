package apiclient

import (
	"github.com/diego-dm-morais/order-manager/gateway/customer"
)

type customerApi struct {
	customer.ICustomerApi
}

func (p *customerApi) FindByCustomer(customerID string) (customer.CustomerOutPutMapper, error) {
	return customer.CustomerOutPutMapper{
		Name:                   "Diego Morais",
		Telephone:              "+55 015 19 9 9999-9999",
		IdentificationDocument: "000.000.000-00",
	}, nil
}
