package routers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Constani/main/repos"
	"github.com/Constani/main/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var constaniCol *mongo.Collection = utils.GetCollection(utils.DB, "constani")
var validate = validator.New()

func Getanim(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Eponime := repos.Comments{
		Content:      "sa",
		AuthorID:     760499240966684683,
		AuthorName:   "Atlas",
		AuthorAvatar: "sa",
	}

	comments := []repos.Comments{Eponime}

	epo1 := repos.Episodes{
		EpisodeNumber:    12,
		EpisodeName:      "sa",
		EpisodesDuration: "asdsadasd",
		Likes:            12,
		VideoURL:         "http:",
		Comments:         comments,
	}

	episodes := []repos.Episodes{
		epo1,
	}
	newAnime := repos.Anime{
		Serie:     "sadasdad",
		TotalLike: 12,
		Episodes:  episodes,
	}
	result, err := constaniCol.InsertOne(ctx, newAnime)
	if err != nil {
		fmt.Printf("Hata var")
	}
	return c.Status(http.StatusCreated).JSON(repos.AnimeResponse{Status: http.StatusCreated, Message: "Sucsess", Data: &fiber.Map{"data": result}})
}
