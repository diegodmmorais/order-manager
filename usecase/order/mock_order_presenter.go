package usecase

import (
	"github.com/stretchr/testify/mock"
)

type MockOrderPresenter struct {
	mock.Mock
}

func (mock *MockOrderPresenter) Success(orderInput OrderSuccessInputData) (*OrderResponse, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*OrderResponse), args.Error(1)
}
