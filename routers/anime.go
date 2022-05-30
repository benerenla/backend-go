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
	return c.Status(http.StatusAccepted).JSON(utils.GetAllData())
}
func CreateAnim(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var anime repos.Anime
	defer cancel()

	if err := c.BodyParser(&anime); err != nil {
		return c.Status(http.StatusBadRequest).JSON(repos.AnimeResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if valodationerr := validate.Struct(&anime); valodationerr != nil {
		return c.Status(http.StatusBadRequest).JSON(repos.AnimeResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": valodationerr.Error()}})
	}
	newAnime := repos.Anime{
		Id:        anime.Id,
		Serie:     anime.Serie,
		TotalLike: anime.TotalLike,
		Episodes:  anime.Episodes,
	}
	result, err := constaniCol.InsertOne(ctx, newAnime)
	if err != nil {
		fmt.Println("Hata var", err.Error())
	}
	return c.Status(http.StatusCreated).JSON(repos.AnimeResponse{Status: http.StatusCreated, Message: "Created", Data: &fiber.Map{"data": result}})
}
func GetAnimeByName(c *fiber.Ctx) error {
	return c.Status(http.StatusAccepted).JSON(utils.GetAnimeById(c.Params("Id")))
}
