package gateway

import order "github.com/diego-dm-morais/order-manager/usecase/order"

func CreateOrderGateway(orderRepository IOrderRepository) order.IOrderGateway {
	return &orderGateway{orderRepository: orderRepository}
}
