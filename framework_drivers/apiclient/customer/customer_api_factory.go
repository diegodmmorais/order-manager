package apiclient

import (
	customer "github.com/diego-dm-morais/order-manager/interface_adapters/gateway/customer"
)

func CreateCustomerApi() customer.ICustomerApi {
	return &customerApi{}
}
