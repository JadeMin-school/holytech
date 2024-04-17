package main

import (
	fiber "github.com/gofiber/fiber/v2"
)
import (
	foodMenu "holytech/foodMenu"
)

var (
	app *fiber.App
)

func init() {
	app = fiber.New()
}



func main() {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(foodMenu.GetToday())
	})

	app.Listen(":3000")
}
