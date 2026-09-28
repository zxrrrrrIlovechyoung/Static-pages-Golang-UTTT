package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8091"
	}

	app.Get("/hello", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Use("/", static.New("./frontend/public"))

	log.Fatal(app.Listen(":" + port))
}
