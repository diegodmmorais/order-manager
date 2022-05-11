package product

import "fmt"

type MockProductGatewayErro struct {
	IProductGateway
}

func (p MockProductGatewayErro) FindByProduct(productID string) (*ProductResponse, error) {
	return nil, fmt.Errorf("produto não encontrado")
}
