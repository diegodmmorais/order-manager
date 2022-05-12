package presenter

import (
	"testing"

	"github.com/diego-dm-morais/order-manager/usecase/order"
	"github.com/stretchr/testify/assert"
)

func Test_imprimir_mensagem_sucesso(t *testing.T) {
	t.Parallel()
	var preseter order.IOrderPresenter = CreateOrderPresenter()
	var orderData order.OrderSuccessInputData = order.OrderSuccessInputData{OrderID: "#1", CustomerName: "Diego Morais"}
	response, err := preseter.Success(orderData)

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
	assert.Nil(t, err)
}

func Test_imprimir_mensagem_sucesso_sem_nome_do_cliente(t *testing.T) {
	t.Parallel()
	var preseter order.IOrderPresenter = CreateOrderPresenter()
	var orderData order.OrderSuccessInputData = order.OrderSuccessInputData{OrderID: "#1"}
	response, err := preseter.Success(orderData)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.NotEmpty(t, err.Error())
	assert.Equal(t, "nome do cliente não encontrado", err.Error())

}

func Test_imprimir_mensagem_sucesso_sem_id_do_pedido(t *testing.T) {
	t.Parallel()
	var preseter order.IOrderPresenter = CreateOrderPresenter()
	var orderData order.OrderSuccessInputData = order.OrderSuccessInputData{CustomerName: "Diego Morais"}
	response, err := preseter.Success(orderData)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.NotEmpty(t, err.Error())
	assert.Equal(t, "id do pedido não encontrado", err.Error())

}
