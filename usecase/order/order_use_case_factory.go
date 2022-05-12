package order

import (
	"github.com/diego-dm-morais/order-manager/usecase/customer"
	"github.com/diego-dm-morais/order-manager/usecase/product"
)

func CreateOrderInputBoundary(
	orderOutputBoundary IOrderPresenter,
	productGateway product.IProductGateway,
	customerGateway customer.ICustomerGateway,
	orderGateway IOrderGateway) IOrderUseCase {
	return &orderUseCase{
		orderOutputBoundary: orderOutputBoundary,
		productGateway:      productGateway,
		customerGateway:     customerGateway,
		orderGateway:        orderGateway,
	}
}
