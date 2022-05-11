package product

type IProductApi interface {
	FindByProduct(productID string) (ProductResponseMapper, error)
}
