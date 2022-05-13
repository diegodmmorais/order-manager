package gateway

import product "github.com/diego-dm-morais/order-manager/usecase/product"

type productGateway struct {
	productApi IProductApi
}

func (p productGateway) FindByProduct(productID string) (product.ProductOutputData, error) {
	productResponse, err := p.productApi.FindByProduct(productID)
	return product.ProductOutputData{
		Nome:                  productResponse.Nome,
		Price:                 productResponse.Price,
		EstoqueEstaDisponivel: productResponse.EstoqueEstaDisponivel,
	}, err
}
