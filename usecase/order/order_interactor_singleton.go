package order

import (
	"sync"

	"github.com/diego-dm-morais/order-manager/usecase/customer"
	"github.com/diego-dm-morais/order-manager/usecase/product"
)

var lock = &sync.Mutex{}

var sigleOrderInteractor *orderInteractor

func CreateOrderInputBoundary(
	orderOutputBoundary IOrderOutputBoundary,
	productGateway product.IProductGateway,
	customerGateway customer.ICustomerGateway,
	orderGateway IOrderGateway) IOrderInputBoundary {
	if sigleOrderInteractor == nil {
		lock.Lock()
		defer lock.Unlock()
		if sigleOrderInteractor == nil {
			sigleOrderInteractor = &orderInteractor{
				orderOutputBoundary: orderOutputBoundary,
				productGateway:      productGateway,
				customerGateway:     customerGateway,
				orderGateway:        orderGateway,
			}
		}
	}
	return sigleOrderInteractor
}
