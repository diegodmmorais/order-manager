package gateway

import (
	customer "github.com/diego-dm-morais/order-manager/usecase/customer"
)

func CreateCustomerGateway(customerApi ICustomerApi) customer.ICustomerGateway {
	return &customerGateway{customerApi: customerApi}
}
