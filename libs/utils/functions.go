package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Constani/main/libs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Get Mongo URL in .env

// database.go, Get collection and client
var col *mongo.Collection = libs.GetCollection(libs.DB, "constani")

// GetAlldata = It provides all anime data to us.
func GetAllData() []libs.Anime {
	var results []libs.Anime

	findOptions := options.Find()
	findOptions.SetLimit(100)
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	cur, err := col.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		fmt.Println("Hata var ", err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result libs.Anime
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}
	return results
}

// GetAnimebyId = <utils>.GetAnimeById("animeId")
func GetAnimeById(ıd string) libs.Anime {
	var result libs.Anime
	filter := bson.D{{"Id", ıd}}
	err := col.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("Hata bulunmaktadır ", err)
	}
	return result
}
