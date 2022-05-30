package libs

import (
	"github.com/gofiber/fiber/v2"
)

// Anime Response
type AnimeResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

//Episodes
type Episodes struct {
	EpisodeNumber    int        `json:"episodeNumber"`
	EpisodeName      string     `json:"episodeName"`
	EpisodesDuration string     `json:"episodesDuration"`
	Likes            int        `json:"likes"`
	Avatar           string     `json"avatar"`
	Description      string     `json:"description"`
	VideoURL         string     `json:"videoURL"`
	Comments         []Comments `json:"Comments"`
}

// Comments
type Comments struct {
	Content      string `json:"content"`
	AuthorID     int    `json:"authorID"`
	AuthorName   string `json:"authorName"`
	AuthorAvatar string `json:"authorAvatar"`
}

// Anime Main
type Anime struct {
	Id        string     `json:"Id,omitempty" bson:"Id" validate:"required"`
	Serie     string     `json:"name,omitempty" validate:"required"`
	Avatar    string     `json:"avatar, omitempty" validate:"required"`
	Banner    string     `json:"banner, omitempty" validate:"required"`
	TotalLike int        `json:"totalLike,omitempty"`
	Episodes  []Episodes `json:"episodes,omitempty" validate:"required"`
}
