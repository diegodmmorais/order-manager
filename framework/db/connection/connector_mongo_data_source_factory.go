package connection

import "github.com/diego-dm-morais/order-manager/repository"

func CreateConnectorMongoDataSource() repository.IConnectorMongoDataSource {
	return &connectorMongoDataSource{}
}
