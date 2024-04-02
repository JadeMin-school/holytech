package main

import (
	foodMenu "holytech/foodMenu"
)
import (
	fiber "github.com/gofiber/fiber/v2"
)

var (
	app *fiber.App
)



func init() {
	app = fiber.New()
}

func main() {
	menu := foodMenu.GetToday()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(menu)
	})

	app.Listen(":3000")
}
