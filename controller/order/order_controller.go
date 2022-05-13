package order

import "github.com/diego-dm-morais/order-manager/usecase/order"

type orderController struct {
	orderUseCase order.IOrderUseCase
}

func (c orderController) Save(request order.OrderRequest) (*order.OrderResponse, error) {
	return c.orderUseCase.Save(request)
}
