package utils

import (
	"os"

	"github.com/Constani/main/logger"
	"github.com/joho/godotenv"
)

func EnvMongo() string {
	err := godotenv.Load()
	if err != nil {
		logger.Logger.FatalF(".env dosyası yuklenmedi %s", err)
	}

	return os.Getenv("MONGODB_URL")
}
