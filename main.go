package main

import (
	"os"

	"github.com/Constani/main/routers"
	"github.com/Constani/main/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	os.Setenv("LOG_FILE", "log.log")
	app.Get("/", routers.GetMain)
	utils.Connect()

	api := app.Group("/api")

	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})
	v1.Get("/anime", routers.Getanim)
	v1.Post("/create", routers.CreateAnim)
	v1.Get("/get/:Id", routers.GetAnimeByName)

	app.Listen(":3000")
}
