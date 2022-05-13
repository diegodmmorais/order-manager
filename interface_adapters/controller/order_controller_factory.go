package controller

import order "github.com/diego-dm-morais/order-manager/application_business_rules/usecase/order"

func CreateOrderController(orderUseCase order.IOrderUseCase) IOrderController {
	return orderController{orderUseCase: orderUseCase}
}
