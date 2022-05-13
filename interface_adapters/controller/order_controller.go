package controller

import order "github.com/diego-dm-morais/order-manager/application_business_rules/usecase/order"

type orderController struct {
	orderUseCase order.IOrderUseCase
}

func (c orderController) Save(request order.OrderRequest) (*order.OrderResponse, error) {
	return c.orderUseCase.Save(request)
}
