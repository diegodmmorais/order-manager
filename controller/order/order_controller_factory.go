package order

import "github.com/diego-dm-morais/order-manager/usecase/order"

func CreateOrderController(orderUseCase order.IOrderUseCase) IOrderController {
	return orderController{orderUseCase: orderUseCase}
}
