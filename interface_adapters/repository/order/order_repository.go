package repository

import (
	"log"

	order "github.com/diego-dm-morais/order-manager/interface_adapters/gateway/order"
	repository "github.com/diego-dm-morais/order-manager/interface_adapters/repository"
	// "github.com/google/uuid"
)

type orderRepository struct {
	order.IOrderRepository
	connectorDataSource repository.IConnectorMongoDataSource
	//connectorDataSource repository.IConnectorPostgresDataSource
}

const DATA_BASE_LABSIT, TABLE_ORDERS = "labsit", "orders"

func (o orderRepository) Save(orderMapper order.OrderInputMapper) (string, error) {
	log.Println("Salvando dados no banco mongo")

	//sql := `
	//	INSERT INTO orders (id, identification_document, freight, total)
	//	VALUES ($1, $2, $3, $4)
	//`
	//o.connectorDataSource.Connect()
	//defer o.connectorDataSource.Disconnect()
	//id, err := o.connectorDataSource.Save(sql, uuid.New().String(), orderMapper.IdentificationDocument, orderMapper.Freight, orderMapper.Total)

	client, _ := o.connectorDataSource.Connect()
	defer o.connectorDataSource.Disconnect(client)
	collection := o.connectorDataSource.DataSource(client, DATA_BASE_LABSIT, TABLE_ORDERS)
	id, err := o.connectorDataSource.Save(collection, orderMapper)
	return id, err
}
