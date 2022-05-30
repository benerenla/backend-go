package routers

import "github.com/gofiber/fiber/v2"

func Handle(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})
	v1.Get("/anime", Getanim)
	v1.Post("/create", CreateAnim)
	v1.Get("/get/:Id", GetAnimeById)
}
