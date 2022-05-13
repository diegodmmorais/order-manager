package controller

import (
	"testing"

	address "github.com/diego-dm-morais/order-manager/usecase/address"
	item "github.com/diego-dm-morais/order-manager/usecase/item"
	order "github.com/diego-dm-morais/order-manager/usecase/order"
	"github.com/stretchr/testify/assert"
)

func Test_salvar_pedido(t *testing.T) {
	t.Parallel()
	ordeUseCase := new(MockOrdeUseCase)
	orderController := CreateOrderController(ordeUseCase)

	ordeUseCase.On("Save").Return(&order.OrderResponse{
		ID:           "#1",
		Message:      "pedido salvo com sucesso",
		CustomerName: "Diego Morais",
	}, nil)

	response, _ := orderController.Save(order.OrderRequest{
		CustomerID: "3",
		Freight:    55.6,
		Items: []item.ItemRequest{
			{
				ProductID: "1",
				Amount:    2,
			},
		},
		ShippingAddress: address.ShippingAddressRequest{
			Street:  "Rua teste",
			Number:  "10",
			Zipcode: "750989899",
			City:    "Goiania",
		},
	})

	assert.NotNil(t, response)
	assert.NotNil(t, response.ID)
	assert.NotEmpty(t, response.ID)
	assert.Equal(t, "#1", response.ID)
	assert.NotNil(t, response.Message)
	assert.NotEmpty(t, response.Message)
	assert.Equal(t, "pedido salvo com sucesso", response.Message)
	assert.NotNil(t, response.CustomerName)
	assert.NotEmpty(t, response.CustomerName)
	assert.Equal(t, "Diego Morais", response.CustomerName)

}
