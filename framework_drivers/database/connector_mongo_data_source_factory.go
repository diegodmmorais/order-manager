package database

import "github.com/diego-dm-morais/order-manager/interface_adapters/repository"

func CreateConnectorMongoDataSource() repository.IConnectorMongoDataSource {
	return &connectorMongoDataSource{}
}
