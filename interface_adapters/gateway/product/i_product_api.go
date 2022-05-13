package gateway

type IProductApi interface {
	FindByProduct(productID string) (ProductOutputMapper, error)
}
