package controller

import (
	order "github.com/diego-dm-morais/order-manager/usecase/order"
	"github.com/stretchr/testify/mock"
)

type MockOrdeUseCase struct {
	mock.Mock
}

func (mock MockOrdeUseCase) Save(request order.OrderRequest) (*order.OrderResponse, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*order.OrderResponse), args.Error(1)
}
