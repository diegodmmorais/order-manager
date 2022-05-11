package customer

import (
	"sync"

	"github.com/diego-dm-morais/order-manager/usecase/customer"
)

var lock = &sync.Mutex{}

var singletonInstance *customerGateway

func CreateCustomerGateway(customerApi ICustomerApi) customer.ICustomerGateway {
	if singletonInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singletonInstance == nil {
			singletonInstance = &customerGateway{customerApi: customerApi}
		}
	}
	return singletonInstance
}
