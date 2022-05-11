package product

type MockProductApi struct {
	IProductApi
}

func (p MockProductApi) FindByProduct(productID string) (ProductResponseMapper, error) {
	return ProductResponseMapper{
		Nome:                  "Macbook 15 pro",
		Price:                 17500,
		EstoqueEstaDisponivel: true,
	}, nil
}
