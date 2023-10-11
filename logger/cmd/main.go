package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

type Config struct {
	Models Models
}

func main() {

	// connect Mongo
	client, err := connectToMongo()

	// create a context in order to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// close connection
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		Models: New(client),
	}

	// listen gRPC
	go app.gRPCListen()

}
