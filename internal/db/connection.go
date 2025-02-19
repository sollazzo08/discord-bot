package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func ConnectToMongoDB() {
	clientOption, err := mongo.Connect(options.Client().ApplyURI(""))
	if err != nil {
		log.Fatal(err)
	 }

	 defer func() {
		if err = clientOption.Disconnect(ctx); err != nil {
				panic(err)
		}
	}()
}




