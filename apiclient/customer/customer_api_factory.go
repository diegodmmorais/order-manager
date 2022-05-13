package apiclient

import (
	customer "github.com/diego-dm-morais/order-manager/gateway/customer"
)

func CreateCustomerApi() customer.ICustomerApi {
	return &customerApi{}
}
