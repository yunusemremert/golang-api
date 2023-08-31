package configs

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(EnvMongoDBURL()))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err := client.Database("books").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client
}

var MongoDB = ConnectMongoDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("books").Collection(collectionName)
}
