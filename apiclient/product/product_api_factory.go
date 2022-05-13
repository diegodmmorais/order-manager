package apiclient

import "github.com/diego-dm-morais/order-manager/gateway/product"

func CreateProductApi() product.IProductApi {
	return &productApi{}
}
