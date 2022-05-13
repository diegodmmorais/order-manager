package repository

import (
	order "github.com/diego-dm-morais/order-manager/interface_adapters/gateway/order"
	repository "github.com/diego-dm-morais/order-manager/interface_adapters/repository"
)

func CreateOrderRepository(connectorMongoDataSource repository.IConnectorMongoDataSource) order.IOrderRepository {
	return &orderRepository{connectorMongoDataSource: connectorMongoDataSource}
}
