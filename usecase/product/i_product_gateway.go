package product

type IProductGateway interface {
	FindByProduct(productID string) (*ProductOutputData, error)
}
