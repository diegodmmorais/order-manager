package database

import "github.com/diego-dm-morais/order-manager/interface_adapters/repository"

func CreateConnectorPostgresDataSource() repository.IConnectorPostgresDataSource {
	return &connectorPostgresDataSource{}
}
