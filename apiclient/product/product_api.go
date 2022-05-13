package apiclient

import (
	product "github.com/diego-dm-morais/order-manager/gateway/product"
)

type productApi struct {
	product.IProductApi
}

func (p *productApi) FindByProduct(productID string) (product.ProductOutputMapper, error) {
	return product.ProductOutputMapper{
		Nome:                  "MacBook 15 pro",
		Price:                 17500.0,
		EstoqueEstaDisponivel: true,
	}, nil
}
