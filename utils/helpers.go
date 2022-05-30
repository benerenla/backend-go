package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Constani/main/logger"
	"github.com/Constani/main/repos"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func EnvMongo() string {
	err := godotenv.Load()
	if err != nil {
		logger.Logger.FatalF(".env dosyası yuklenmedi %s", err)
	}

	return os.Getenv("MONGODB_URL")
}

var col *mongo.Collection = GetCollection(DB, "constani")

func GetAllData() []repos.Anime {
	var results []repos.Anime

	findOptions := options.Find()
	findOptions.SetLimit(100)
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	cur, err := col.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		fmt.Println("Hata var ", err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result repos.Anime
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}
	return results
}
func GetAnimeById(ıd string) repos.Anime {
	var result repos.Anime
	filter := bson.D{{"Id", ıd}}
	err := col.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("Hata bulunmaktadır ", err)
	}
	return result
}
