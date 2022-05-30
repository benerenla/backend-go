package libs

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	// User Id example 55221047294284827
	Id string `json:"id, omitempty" bson:"Id" validate: "required"`

	//username Example atlas
	Username string `json:"username, omitempty" validate: "required"`

	// User Account Created At
	CreatedAt string `json:"createdAt, omitempty" validate:"required"`

	// User Avatar URL
	AvatarURL string `json:"avatarURL, omitempty" validate: "required"`

	// BannerURL
	BannerURL string `json:"banner_url, omitempty" validate: "required"`

	// User About Me.
	AboutMe string `json:"description, omitempty" validate: "required"`

	Status string `json:"status, omitempty"`

	// User Flags Admin, Developer, Moderator ...
	Flags []string `json:"flags, omitempty" validate: "required"`

	// User Followers
	Followers string `json:"followers, omitempty" validate: "required"`

	// User Favs
	Favs []string `json:"favs, omitempty" validate: "required"`

	// User Liked Animes
	LikedAnimes []string `json:"likedAnimes, omitempty" validate: "required"`

	// Following
	Following []int `json:"following, omitempty" validate: "required"`

	// Anime Watch Later
	WatchLater []string `json:"watchLater, omitempty" validate: "required"`
}
type UserResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

var constaniCol *mongo.Collection = GetCollection(DB, "users")
var validate = validator.New()

func CreateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user User
	defer cancel()

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if valodationerr := validate.Struct(&user); valodationerr != nil {
		return c.Status(http.StatusBadRequest).JSON(UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": valodationerr.Error()}})
	}
	newUser := User{
		Id:          user.Id,
		Username:    user.Username,
		CreatedAt:   user.CreatedAt,
		AvatarURL:   user.AvatarURL,
		BannerURL:   user.BannerURL,
		AboutMe:     user.AboutMe,
		Status:      user.Status,
		Flags:       user.Flags,
		Followers:   user.Followers,
		Favs:        user.Favs,
		LikedAnimes: user.LikedAnimes,
		Following:   user.Following,
		WatchLater:  user.WatchLater,
	}
	result, err := constaniCol.InsertOne(ctx, newUser)
	if err != nil {
		fmt.Println("Hata var", err.Error())
	}
	return c.Status(http.StatusCreated).JSON(UserResponse{Status: http.StatusCreated, Message: "Created", Data: &fiber.Map{"data": result}})
}
func GetUser(c *fiber.Ctx) error {
	var result User
	filter := bson.D{{"Id", c.Params("Id")}}
	err := constaniCol.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("Hata bulunmaktadır ", err)
	}
	return c.Status(http.StatusAccepted).JSON(result)
}
