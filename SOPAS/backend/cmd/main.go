package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func main() {
	app := fiber.New()
	app.Use("/", static.New("./frontend/public"))
	app.Get("/hello", func(c fiber.Ctx) error {
		return c.SendString("hello world")
	})
	log.Fatal(app.Listen(":8093"))

}
