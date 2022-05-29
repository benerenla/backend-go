package routers

import (
	"github.com/gofiber/fiber/v2"
)

func GetMain(c *fiber.Ctx) error {
	return c.SendString("Hello, World! ")
}
