package apiclient

import product "github.com/diego-dm-morais/order-manager/gateway/product"

func CreateProductApi() product.IProductApi {
	return &productApi{}
}
