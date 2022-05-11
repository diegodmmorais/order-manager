package presenter

import (
	"fmt"
	"strings"

	"github.com/diego-dm-morais/order-manager/usecase/order"
)

type orderPresenter struct {
	order.IOrderOutputBoundary
}

func (p orderPresenter) Success(orderInput order.OrderInputData) (*order.OrderResponse, error) {
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
