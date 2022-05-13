package gateway

import order "github.com/diego-dm-morais/order-manager/application_business_rules/usecase/order"

func CreateOrderGateway(orderRepository IOrderRepository) order.IOrderGateway {
	return &orderGateway{orderRepository: orderRepository}
}
