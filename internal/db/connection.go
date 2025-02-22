package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectToMongoDB(URI string) error {
	uri := URI
	docs := "www.mongodb.com/docs/drivers/go/current/"
	if uri == "" {
		return fmt.Errorf("set your 'MONGODB_URI' environment variable. See: %s", docs)
	}
	client, err := mongo.Connect(options.Client().
		ApplyURI(uri))
	if err != nil {
		return fmt.Errorf("error connecting to MongoDB: %w", err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Println("Error disconnecting MongoDB:", err)
		}
	}()

	return nil
}
