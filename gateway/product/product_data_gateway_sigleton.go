package product

import (
	"sync"

	"github.com/diego-dm-morais/order-manager/usecase/product"
)

var lock = &sync.Mutex{}

var singletonInstance *productGateway

func CreateProductGateway(productApi IProductApi) product.IProductGateway {
	if singletonInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singletonInstance == nil {
			singletonInstance = &productGateway{productApi: productApi}
		}
	}
	return singletonInstance
}
