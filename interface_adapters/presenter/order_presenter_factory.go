package presenter

import (
	order "github.com/diego-dm-morais/order-manager/application_business_rules/usecase/order"
)

func CreateOrderPresenter() order.IOrderPresenter {
	return &orderPresenter{}
}
