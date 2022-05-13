package presenter

import (
	"fmt"
	"log"
	"strings"

	order "github.com/diego-dm-morais/order-manager/application_business_rules/usecase/order"
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
		log.Println("Retornando operação executada com sucesso")
		return &order.OrderResponse{
			ID:           orderInput.OrderID,
			CustomerName: orderInput.CustomerName,
			Message:      "pedido salvo com sucesso",
		}, nil
	}
}
