package usecase

import (
	customer "github.com/diego-dm-morais/order-manager/application_business_rules/usecase/customer"
	product "github.com/diego-dm-morais/order-manager/application_business_rules/usecase/product"
)

func CreateOrderUseCase(
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
