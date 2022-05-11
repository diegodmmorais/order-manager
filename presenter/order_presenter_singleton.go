package presenter

import (
	"sync"

	"github.com/diego-dm-morais/order-manager/usecase/order"
)

var lock = &sync.Mutex{}

var sigletonInstance *orderPresenter

func CreateOrderPresenter() order.IOrderOutputBoundary {
	if sigletonInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if sigletonInstance == nil {
			sigletonInstance = &orderPresenter{}
		}
	}
	return sigletonInstance
}
