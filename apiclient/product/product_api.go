package apiclient

import (
	"github.com/diego-dm-morais/order-manager/gateway/product"
)

type productApi struct {
	product.IProductApi
}

func (p *productApi) FindByProduct(productID string) (product.ProductOutputMapper, error) {
	return product.ProductOutputMapper{
		Nome:                  "MacBook 15 pro",
		Price:                 float32(17500.0),
		EstoqueEstaDisponivel: true,
	}, nil
}