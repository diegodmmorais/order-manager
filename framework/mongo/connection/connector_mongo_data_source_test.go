package connection

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id    primitive.ObjectID `bson:"_id"`
	Total float32            `bson:"total"`
	Itens []Item             `bson:"itens"`
}

type Item struct {
	ID         primitive.ObjectID `bson:"_id"`
	Nome       string             `bson:"nome"`
	Quantidade uint32             `bson:"quantidade"`
	Preco      float32            `bson:"preco"`
}

func TestConectarComUmBanco(t *testing.T) {
	t.Parallel()
	dataSource := CreateConnectorMongoDataSource()

	client, _ := dataSource.Connect()

	collection := dataSource.DataSource(client, "lw-orders", "orders")

	collection.InsertOne(context.TODO(), &Order{
		Id:    primitive.NewObjectID(),
		Total: float32(1000.0),
		Itens: []Item{
			{
				ID:         primitive.NewObjectID(),
				Nome:       "Macbook",
				Quantidade: 1,
				Preco:      float32(1000.0),
			},
		},
	})
}
