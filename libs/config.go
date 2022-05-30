package libs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongo() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Hata var", err)
	}

	return os.Getenv("MONGODB_URL")
}

func GetWeebhook() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Hata var", err)
	}
	return os.Getenv("WEBHOOK_URL")
}

func GetCollectionName() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Hata var", err)
	}
	return os.Getenv("COLLECTION_NAME")
}
