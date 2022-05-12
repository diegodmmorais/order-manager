package repository

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Task struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Text      string             `bson:"text"`
	Completed bool               `bson:"completed"`
}

var collection *mongo.Collection
var client *mongo.Client
var ctx = context.TODO()

func inicializar() {
	clientOptions := options.Client().ApplyURI("mongodb://admin:admin123@127.0.0.1:27017")
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("tasker").Collection("tasks")
}

func disconect() {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func createTask(task *Task) (string, error) {
	teste, err := collection.InsertOne(ctx, task)
	id := teste.InsertedID.(primitive.ObjectID).Hex()
	return id, err
}

func Test_yrdyr(t *testing.T) {
	inicializar()
	createTask(&Task{
		ID:        primitive.NewObjectID(),
		Text:      "olá",
		Completed: true,
	})
	disconect()
}
