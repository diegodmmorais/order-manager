package repository

type IConnectorPostgresDataSource interface {
	Connect() error
	Disconnect() (bool, error)
	Save(sql string, args ...any) (string, error)
}
