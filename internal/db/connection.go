package db

import (
	"fmt"
	"context"


	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func ConnectToMongoDB() {
	client, err := mongo.Connect(options.Client().
		ApplyURI(""))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

}





