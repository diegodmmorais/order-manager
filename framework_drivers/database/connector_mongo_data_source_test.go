package database

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id    primitive.ObjectID `bson:"_id"`
	Total float64            `bson:"total"`
	Itens []Item             `bson:"itens"`
}

type Item struct {
	ID         primitive.ObjectID `bson:"_id"`
	Nome       string             `bson:"nome"`
	Quantidade uint32             `bson:"quantidade"`
	Preco      float64            `bson:"preco"`
}

func TestConectarComUmBanco(t *testing.T) {
	t.Parallel()
	dataSource := CreateConnectorMongoDataSource()

	client, _ := dataSource.Connect()

	collection := dataSource.DataSource(client, "labsit", "orders")

	collection.InsertOne(context.TODO(), &Order{
		Id:    primitive.NewObjectID(),
		Total: 1000.23,
		Itens: []Item{
			{
				ID:         primitive.NewObjectID(),
				Nome:       "Macbook",
				Quantidade: 1,
				Preco:      1000.23,
			},
		},
	})
}
