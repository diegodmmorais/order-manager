package order

import (
	"fmt"
	"testing"

	"github.com/diego-dm-morais/order-manager/usecase/address"
	"github.com/diego-dm-morais/order-manager/usecase/item"
	"github.com/diego-dm-morais/order-manager/usecase/order"
	"github.com/stretchr/testify/assert"
)

func Test_salvar_pedio(t *testing.T) {
	t.Parallel()
	orderRepository := new(MockOrderRepository)

	orderGateway := CreateOrderGateway(orderRepository)

	orderRepository.On("Save").Return("#1", nil)

	id, err := orderGateway.Save(order.OrderInputData{
		IdentificationDocument: "999.999.999-00",
		Freight:                float32(55.0),
		Total:                  17555.0,
		Itens: []item.ItemInputData{
			{
				ProductName: "Macbook pro 15",
				Price:       float32(17500.0),
				Amount:      1,
			},
		},
		ShippingAddress: address.AddressInputData{
			Street:  "Rua new street",
			Number:  "8987",
			Zipcode: "89898-09",
			City:    "São Paulo - SP",
		},
	})

	assert.NotEmpty(t, id)
	assert.Equal(t, "#1", id)
	assert.Nil(t, err)

}

func Test_erro_ao_salvar_pedio(t *testing.T) {
	t.Parallel()
	orderRepository := new(MockOrderRepository)

	orderGateway := CreateOrderGateway(orderRepository)

	orderRepository.On("Save").Return("", fmt.Errorf("Erro inesperado ao salvar o pedido"))

	id, err := orderGateway.Save(order.OrderInputData{
		IdentificationDocument: "999.999.999-00",
		Freight:                float32(55.0),
		Total:                  17555.0,
		Itens: []item.ItemInputData{
			{
				ProductName: "Macbook pro 15",
				Price:       float32(17500.0),
				Amount:      1,
			},
		},
		ShippingAddress: address.AddressInputData{
			Street:  "Rua new street",
			Number:  "8987",
			Zipcode: "89898-09",
			City:    "São Paulo - SP",
		},
	})

	assert.Empty(t, id)
	assert.NotNil(t, err)
	assert.NotEmpty(t, err.Error())
	assert.Equal(t, "Erro inesperado ao salvar o pedido", err.Error())

}
