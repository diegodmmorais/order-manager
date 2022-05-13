package repository

import (
	order "github.com/diego-dm-morais/order-manager/gateway/order"
	repository "github.com/diego-dm-morais/order-manager/repository"
)

func CreateOrderRepository(connectorMongoDataSource repository.IConnectorMongoDataSource) order.IOrderRepository {
	return &orderRepository{connectorMongoDataSource: connectorMongoDataSource}
}
