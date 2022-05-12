package order

import (
	"github.com/diego-dm-morais/order-manager/gateway/order"
	"github.com/diego-dm-morais/order-manager/repository"
)

func CreateOrderRepository(connectorMongoDataSource repository.IConnectorMongoDataSource) order.IOrderRepository {
	return &orderRepository{connectorMongoDataSource: connectorMongoDataSource}
}
