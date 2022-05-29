package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Constani/main/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongo()))
	if err != nil {
		logger.Logger.FatalF("mongo db bağlanmadı hata %s", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		logger.Logger.FatalF("mongo db bağlanmadı hata %s", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Hata")
	}
	fmt.Printf("Connected Mongodb")
	return client
}

var DB *mongo.Client = Connect()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("constani").Collection(collectionName)
	return collection
}
