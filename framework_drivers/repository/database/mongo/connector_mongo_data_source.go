package database

import (
	"context"
	"log"

	"github.com/diego-dm-morais/order-manager/interface_adapters/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joho/godotenv"
	"os"
)

type connectorMongoDataSource struct {
	repository.IConnectorMongoDataSource
}

var clientOptions *options.ClientOptions

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
	clientOptions = options.Client().ApplyURI(os.Getenv("MONGO_DATA_BASE_URL"))
}

func (c *connectorMongoDataSource) Connect() (*mongo.Client, error) {
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Connection to MongoDB opened.")
	return client, nil
}

func (c *connectorMongoDataSource) DataSource(client *mongo.Client, database, collection string) *mongo.Collection {
	return client.Database(database).Collection(collection)
}

func (c *connectorMongoDataSource) Disconnect(client *mongo.Client) (bool, error) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Println(err)
		return false, err
	}
	log.Println("Connection to MongoDB closed.")
	return true, nil
}

func (c *connectorMongoDataSource) Save(collection *mongo.Collection, date interface{}) (string, error) {
	teste, err := collection.InsertOne(context.TODO(), date)
	return teste.InsertedID.(primitive.ObjectID).Hex(), err
}
