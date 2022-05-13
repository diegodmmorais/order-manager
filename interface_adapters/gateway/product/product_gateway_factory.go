package gateway

import (
	product "github.com/diego-dm-morais/order-manager/application_business_rules/usecase/product"
)

func CreateProductGateway(productApi IProductApi) product.IProductGateway {
	return productGateway{productApi: productApi}
}
