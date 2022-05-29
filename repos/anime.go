package repos

import (
	"github.com/gofiber/fiber/v2"
)

type AnimeResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

type Episodes struct {
	EpisodeNumber    int        `json:"episodeNumber"`
	EpisodeName      string     `json:"episodeName"`
	EpisodesDuration string     `json:"episodesDuration"`
	Likes            int        `json:"likes"`
	VideoURL         string     `json:"videoURL"`
	Comments         []Comments `json:"Comments"`
}
type Comments struct {
	Content      string `json:"content"`
	AuthorID     int    `json:"authorID"`
	AuthorName   string `json:"authorName"`
	AuthorAvatar string `json:"authorAvatar"`
}
type Anime struct {
	Serie     string     `json:"name,omitempty" validate:"required"`
	TotalLike int        `json:"totalLike,omitempty"`
	Episodes  []Episodes `json:"episodes,omitempty" validate:"required"`
}
