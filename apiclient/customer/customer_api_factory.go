package apiclient

import (
	"github.com/diego-dm-morais/order-manager/gateway/customer"
)

func CreateCustomerApi() customer.ICustomerApi {
	return &customerApi{}
}
