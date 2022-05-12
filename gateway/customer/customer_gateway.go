package customer

import "github.com/diego-dm-morais/order-manager/usecase/customer"

type customerGateway struct {
	customer.ICustomerGateway
	customerApi ICustomerApi
}

func (c customerGateway) FindByCustomer(customerID string) (customer.CustomerOutputData, error) {
	response, err := c.customerApi.FindByCustomer(customerID)
	return customer.CustomerOutputData{
		IdentificationDocument: response.IdentificationDocument,
		Name:                   response.Name,
		Telephone:              response.Telephone,
	}, err
}
