package presenter

import (
	"fmt"
	"strings"

	order "github.com/diego-dm-morais/order-manager/usecase/order"
)

type orderPresenter struct {
	order.IOrderPresenter
}

func (p orderPresenter) Success(orderInput order.OrderSuccessInputData) (*order.OrderResponse, error) {
	switch {
	case orderInput.CustomerName == "" || strings.TrimSpace(orderInput.CustomerName) == "":
		return nil, fmt.Errorf("nome do cliente não encontrado")
	case orderInput.OrderID == "" || strings.TrimSpace(orderInput.OrderID) == "":
		return nil, fmt.Errorf("id do pedido não encontrado")
	default:
		return &order.OrderResponse{
			ID:           orderInput.OrderID,
			CustomerName: orderInput.CustomerName,
			Message:      "pedido salvo com sucesso",
		}, nil
	}
}
