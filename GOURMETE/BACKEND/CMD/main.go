package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func main() {
	app := fiber.New()
	app.Use("/", static.New("./FRONTEND/public"))
	app.Get("/hello", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	log.Fatal(app.Listen(":8080"))

}
