package main

import (
	"github.com/Constani/main/libs"
	"github.com/Constani/main/libs/routers"
	"github.com/Constani/main/libs/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:8080, http://localhost:3131, https://127.0.0.1:3131, https://127.0.0.1:8080",
		AllowHeaders:     "Origin, Content-Type, Accept, Accept-Language, Content-Length",
	}))

	libs.Connect()
	utils.SendMessage("Backend başarıyla aktif hale geldi.", "Site Başlatıldı", "https://atlasch.me")
	routers.Handle(app)

	app.Listen(":3000")
}
