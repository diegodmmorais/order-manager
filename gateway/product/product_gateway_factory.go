package product

import (
	"github.com/diego-dm-morais/order-manager/usecase/product"
)

func CreateProductGateway(productApi IProductApi) product.IProductGateway {
	return productGateway{productApi: productApi}
}
