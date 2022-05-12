package product

import (
	"testing"

	"github.com/diego-dm-morais/order-manager/usecase/product"
	"github.com/stretchr/testify/assert"
)

func Test_buscando_um_produto(t *testing.T) {
	t.Parallel()

	mockProductApi := new(MockProductApi)
	var productGateway product.IProductGateway = CreateProductGateway(mockProductApi)

	mockProductApi.On("FindByProduct").Return(ProductOutputMapper{
		Nome:                  "Macbook 15 pro",
		Price:                 float32(17500),
		EstoqueEstaDisponivel: true,
	}, nil)

	response, err := productGateway.FindByProduct("#1")

	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.NotNil(t, response.Nome)
	assert.NotEmpty(t, response.Nome)
	assert.Equal(t, "Macbook 15 pro", response.Nome)

	assert.NotNil(t, response.Price)
	assert.Equal(t, float32(17500), response.Price)

	assert.NotNil(t, response.EstoqueEstaDisponivel)
	assert.Equal(t, true, response.EstoqueEstaDisponivel)
}
