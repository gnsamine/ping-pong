package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type JsonMessage struct {
	Message string `json:"message"`
}

func GetJson(c *fiber.Ctx) error {

	queryValue := c.Query("message")

	MakeUppercase := strings.ToUpper(queryValue)

	return c.SendString(MakeUppercase)
}

func main() {
	app := fiber.New()
	app.Get("/amine", func(c *fiber.Ctx) error {

		queryValue := c.Query("message")

		if queryValue != "" {
			MakeUppercase := strings.ToUpper(queryValue)

			return c.SendString(MakeUppercase)

		} else {
			return c.SendStatus(400)
		}

	})
	log.Fatal(app.Listen(":7070"))
}
