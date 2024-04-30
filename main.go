package main

import (
	fiber "github.com/gofiber/fiber/v2"
	cors "github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	

	app.Get("/today", func(c *fiber.Ctx) error {
		return c.JSON(foodMenu.GetToday())
	})
	app.Get("/week", func(c *fiber.Ctx) error {
		return c.JSON(foodMenu.GetThisWeek())
	})

	app.Listen(":3000")
}