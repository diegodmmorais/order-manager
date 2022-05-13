package controller

import order "github.com/diego-dm-morais/order-manager/application_business_rules/usecase/order"

type IOrderController interface {
	Save(OrderRequest order.OrderRequest) (*order.OrderResponse, error)
}
