package database

import (
	"log"

	"database/sql"

	"os"

	"github.com/diego-dm-morais/order-manager/interface_adapters/repository"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type connectorPostgresDataSource struct {
	repository.IConnectorPostgresDataSource
	clientOptions *sql.DB
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
}

func (c *connectorPostgresDataSource) Connect() error {
	log.Printf("conectando %s", os.Getenv("POSTGRES_DATA_BASE_URL"))
	var err error
	c.clientOptions, err = sql.Open("postgres", os.Getenv("POSTGRES_DATA_BASE_URL"))
	if err != nil {
		log.Println(err)
		return err
	}
	pingErr := c.clientOptions.Ping()
	if pingErr != nil {
		log.Println(pingErr)
		return pingErr
	}
	log.Println("Connection to MongoDB opened.")
	return nil
}

func (c *connectorPostgresDataSource) Disconnect() (bool, error) {
	err := c.clientOptions.Close()
	if err != nil {
		log.Println(err)
		return false, err
	}
	log.Println("Connection to MongoDB closed.")
	return true, nil
}

func (c *connectorPostgresDataSource) Save(sql string, args ...any) (string, error) {
	id := ""
	script := sql + " RETURNING id"
	err := c.clientOptions.QueryRow(script, args...).Scan(&id)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return id, nil
}
