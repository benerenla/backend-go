package routers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Constani/main/libs"
	"github.com/Constani/main/libs/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var constaniCol *mongo.Collection = libs.GetCollection(libs.DB, "constani")
var validate = validator.New()

// Get All Anime.
func Getanim(c *fiber.Ctx) error {
	return c.Status(http.StatusAccepted).JSON(utils.GetAllData())
}

// Post: Create New Anime
func CreateAnim(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var anime libs.Anime
	defer cancel()

	if err := c.BodyParser(&anime); err != nil {
		return c.Status(http.StatusBadRequest).JSON(libs.AnimeResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if valodationerr := validate.Struct(&anime); valodationerr != nil {
		return c.Status(http.StatusBadRequest).JSON(libs.AnimeResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": valodationerr.Error()}})
	}
	newAnime := libs.Anime{
		Id:        anime.Id,
		Serie:     anime.Serie,
		Avatar:    anime.Avatar,
		Banner:    anime.Banner,
		TotalLike: anime.TotalLike,
		Episodes:  anime.Episodes,
	}
	result, err := constaniCol.InsertOne(ctx, newAnime)
	if err != nil {
		fmt.Println("Hata var", err.Error())
	}
	var message = "**" + newAnime.Serie + "** adında yeni anime siteye eklendi!"
	utils.SendMessage(message, "Yeni Anime Eklendi!", "http://atlasch.me")
	return c.Status(http.StatusCreated).JSON(libs.AnimeResponse{Status: http.StatusCreated, Message: "Created", Data: &fiber.Map{"data": result}})
}

// GetAnimeById = api/v1/get/<animeId>
func GetAnimeById(c *fiber.Ctx) error {
	return c.Status(http.StatusAccepted).JSON(utils.GetAnimeById(c.Params("Id")))
}
