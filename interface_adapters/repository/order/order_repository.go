package repository

import (
	order "github.com/diego-dm-morais/order-manager/interface_adapters/gateway/order"
	repository "github.com/diego-dm-morais/order-manager/interface_adapters/repository"
)

type orderRepository struct {
	order.IOrderRepository
	connectorMongoDataSource repository.IConnectorMongoDataSource
}

const DATA_BASE_LABSIT, TABLE_ORDERS = "labsit", "orders"

func (o orderRepository) Save(orderMapper order.OrderInputMapper) (string, error) {
	client, _ := o.connectorMongoDataSource.Connect()
	collection := o.connectorMongoDataSource.DataSource(client, DATA_BASE_LABSIT, TABLE_ORDERS)
	id, err := o.connectorMongoDataSource.Save(collection, orderMapper)
	o.connectorMongoDataSource.Disconnect(client)
	return id, err
}
