package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func Connect() (*mongo.Client, *context.Context, error) {

	ctx := context.TODO()
	clientOption := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))

	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		return nil, nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	return client, &ctx, nil

}

func Disconnect(client *mongo.Client, ctx *context.Context) {
	err := client.Disconnect(*ctx)
	if err != nil {
		log.Fatal(err)
	}
}
