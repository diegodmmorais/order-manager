package usecase

type IProductGateway interface {
	FindByProduct(productID string) (ProductOutputData, error)
}
