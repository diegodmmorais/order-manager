package product

type MockProductGateway struct {
	IProductGateway
}

func (p MockProductGateway) FindByProduct(productID string) ProductResponse {
	return ProductResponse{
		Nome:                  "Macbook 15 pro",
		Price:                 17.0090,
		EstoqueEstaDisponivel: true,
	}
}