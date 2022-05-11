package order

import (
	"testing"

	"github.com/diego-dm-morais/order-manager/usecase/address"
	"github.com/diego-dm-morais/order-manager/usecase/customer"
	"github.com/diego-dm-morais/order-manager/usecase/item"
	"github.com/diego-dm-morais/order-manager/usecase/product"
	"github.com/stretchr/testify/assert"
)

var orderInputBoundary IOrderInputBoundary = CreateOrderInputBoundary(MockOrderOutputBoundary{}, product.MockProductGateway{}, customer.MockCustomerGateway{}, MockOrderGateway{})

func Test_salvar_pedido(t *testing.T) {
	t.Parallel()
	response, erro := orderInputBoundary.Salvar(OrderRequest{
		CustomerID: "1",
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
	assert.NotNil(t, response.CustomerName)
	assert.NotEmpty(t, response.CustomerName)
	assert.Equal(t, "Diego", response.CustomerName)

	assert.NotNil(t, response.Message)
	assert.NotEmpty(t, response.Message)
	assert.Equal(t, "pedido salvo com sucesso", response.Message)

	assert.Nil(t, erro)

}

func Test_salvar_pedido_sem_itens(t *testing.T) {
	t.Parallel()
	response, erro := orderInputBoundary.Salvar(OrderRequest{
		CustomerID: "1",
		Freight:    55.6,
		ShippingAddress: address.ShippingAddressRequest{
			Street:  "Rua teste",
			Number:  "10",
			Zipcode: "750989899",
			City:    "Goiania",
		},
	})

	assert.Nil(t, response)
	assert.NotNil(t, erro)
	assert.NotEmpty(t, erro.Error())
	assert.Equal(t, "itens não foram encontrados", erro.Error())

}

func Test_salvar_pedido_com_erro_ao_persistir(t *testing.T) {
	t.Parallel()
	var orderInputBoundary IOrderInputBoundary = &orderInteractor{
		orderOutputBoundary: MockOrderOutputBoundary{},
		productGateway:      product.MockProductGateway{},
		customerGateway:     customer.MockCustomerGateway{},
		orderGateway:        MockOrderGatewayError{},
	}

	response, erro := orderInputBoundary.Salvar(OrderRequest{
		CustomerID: "1",
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

	assert.Nil(t, response)
	assert.NotNil(t, erro)
	assert.NotEmpty(t, erro.Error())
	assert.Equal(t, "erro ao salvar o produto", erro.Error())

}

func Test_erro_ocorrido_ao_buscar_produtos(t *testing.T) {
	t.Parallel()
	var orderInputBoundary IOrderInputBoundary = &orderInteractor{
		orderOutputBoundary: MockOrderOutputBoundary{},
		productGateway:      product.MockProductGatewayErro{},
		customerGateway:     customer.MockCustomerGateway{},
		orderGateway:        MockOrderGateway{},
	}

	response, erro := orderInputBoundary.Salvar(OrderRequest{
		CustomerID: "1",
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

	assert.Nil(t, response)
	assert.NotNil(t, erro)
	assert.NotEmpty(t, erro.Error())
	assert.Equal(t, "itens não foram encontrados", erro.Error())

}
