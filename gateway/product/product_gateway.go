package product

import "github.com/diego-dm-morais/order-manager/usecase/product"

type productGateway struct {
	product.IProductGateway
	productApi IProductApi
}

func (p productGateway) FindByProduct(productID string) (*product.ProductResponse, error) {
	productResponse, err := p.productApi.FindByProduct(productID)
	return &product.ProductResponse{
		Nome:                  productResponse.Nome,
		Price:                 productResponse.Price,
		EstoqueEstaDisponivel: productResponse.EstoqueEstaDisponivel,
	}, err
}
