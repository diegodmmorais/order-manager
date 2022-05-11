package product

type IProductGateway interface {
	FindByProduct(productID string) ProductResponse
}
