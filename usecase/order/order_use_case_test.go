package usecase

import (
	"fmt"
	"testing"

	address "github.com/diego-dm-morais/order-manager/usecase/address"
	customer "github.com/diego-dm-morais/order-manager/usecase/customer"
	item "github.com/diego-dm-morais/order-manager/usecase/item"
	product "github.com/diego-dm-morais/order-manager/usecase/product"
	"github.com/stretchr/testify/assert"
)

func Test_salvar_pedido(t *testing.T) {
	t.Parallel()
	mockOrderOutput := new(MockOrderPresenter)
	mockIProductGateway := new(product.MockProductGateway)
	mockICustomerGateway := new(customer.MockCustomerGateway)
	mockIOrderGateway := new(MockOrderGateway)
	orderInputBoundary := CreateOrderUseCase(mockOrderOutput, mockIProductGateway, mockICustomerGateway, mockIOrderGateway)

	mockIProductGateway.On("FindByProduct").Return(product.ProductOutputData{
		Nome:                  "Macbook pro 15",
		Price:                 17500.0,
		EstoqueEstaDisponivel: true,
	}, nil)

	mockICustomerGateway.On("FindByCustomer").Return(customer.CustomerOutputData{
		Name:                   "Diego Morais",
		Telephone:              "+55 015 9 9999-9999",
		IdentificationDocument: "000.000.000-00",
	}, nil)

	mockIOrderGateway.On("Save").Return("#1", nil)

	mockOrderOutput.On("Success").Return(&OrderResponse{
		ID:           "#1",
		Message:      "pedido salvo com sucesso",
		CustomerName: "Diego Morais",
	}, nil)

	orderRequest := OrderRequest{
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
	}

	response, erro := orderInputBoundary.Save(orderRequest)

	assert.NotNil(t, response)
	assert.NotNil(t, response.ID)
	assert.NotEmpty(t, response.ID)
	assert.Equal(t, "#1", response.ID)
	assert.NotNil(t, response.CustomerName)
	assert.NotEmpty(t, response.CustomerName)
	assert.Equal(t, "Diego Morais", response.CustomerName)

	assert.NotNil(t, response.Message)
	assert.NotEmpty(t, response.Message)
	assert.Equal(t, "pedido salvo com sucesso", response.Message)

	assert.Nil(t, erro)
}

func Test_salvar_pedido_sem_itens(t *testing.T) {
	t.Parallel()
	mockOrderOutput := new(MockOrderPresenter)
	mockIProductGateway := new(product.MockProductGateway)
	mockICustomerGateway := new(customer.MockCustomerGateway)
	mockIOrderGateway := new(MockOrderGateway)
	orderInputBoundary := CreateOrderUseCase(mockOrderOutput, mockIProductGateway, mockICustomerGateway, mockIOrderGateway)

	mockICustomerGateway.On("FindByCustomer").Return(customer.CustomerOutputData{
		Name:                   "Diego Morais",
		Telephone:              "+55 015 9 9999-9999",
		IdentificationDocument: "000.000.000-00",
	}, nil)

	orderRequest := OrderRequest{
		CustomerID: "3",
		Freight:    55.6,
		ShippingAddress: address.ShippingAddressRequest{
			Street:  "Rua teste",
			Number:  "10",
			Zipcode: "750989899",
			City:    "Goiania",
		},
	}

	response, erro := orderInputBoundary.Save(orderRequest)

	assert.Nil(t, response)
	assert.NotNil(t, erro)
	assert.NotEmpty(t, erro.Error())
	assert.Equal(t, "itens não foram encontrados", erro.Error())

}

func Test_salvar_pedido_erro_ao_salvar_pedido(t *testing.T) {
	t.Parallel()
	mockOrderOutput := new(MockOrderPresenter)
	mockIProductGateway := new(product.MockProductGateway)
	mockICustomerGateway := new(customer.MockCustomerGateway)
	mockIOrderGateway := new(MockOrderGateway)
	orderInputBoundary := CreateOrderUseCase(mockOrderOutput, mockIProductGateway, mockICustomerGateway, mockIOrderGateway)

	mockIProductGateway.On("FindByProduct").Return(product.ProductOutputData{
		Nome:                  "Macbook pro 15",
		Price:                 float64(17500),
		EstoqueEstaDisponivel: true,
	}, nil)

	mockICustomerGateway.On("FindByCustomer").Return(customer.CustomerOutputData{
		Name:                   "Diego Morais",
		Telephone:              "+55 015 9 9999-9999",
		IdentificationDocument: "000.000.000-00",
	}, nil)

	mockIOrderGateway.On("Save").Return("", fmt.Errorf("erro ao salvar o pedido"))

	orderRequest := OrderRequest{
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
	}

	response, erro := orderInputBoundary.Save(orderRequest)

	fmt.Println(response)
	assert.Nil(t, response)
	assert.NotNil(t, erro)
	assert.NotEmpty(t, erro.Error())
	assert.Equal(t, "erro ao salvar o pedido", erro.Error())

}
