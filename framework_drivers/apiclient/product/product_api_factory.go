package apiclient

import product "github.com/diego-dm-morais/order-manager/interface_adapters/gateway/product"

func CreateProductApi() product.IProductApi {
	return &productApi{}
}
