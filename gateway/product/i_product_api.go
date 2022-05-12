package product

type IProductApi interface {
	FindByProduct(productID string) (ProductOutputMapper, error)
}
