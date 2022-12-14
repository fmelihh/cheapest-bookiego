package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var MongoUri = "mongodb://root:root@localhost:27017"

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func connect() (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoUri))
	return client, ctx, cancel, err
}

func ping(client *mongo.Client, ctx context.Context) error {

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Print("connected successfully")
	return nil
}
