package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type JsonMessage struct {
	Message string `json:"text"`
}

func main() {
	app := fiber.New()
	app.Get("/amine", func(c *fiber.Ctx) error {

		queryValue := c.Query("text")

		if queryValue != "" {

			var Jsonmes JsonMessage

			MakeUppercase := strings.ToUpper(queryValue)

			Jsonmes.Message = MakeUppercase

			return c.JSON(MakeUppercase)

		} else {
			return c.SendStatus(400)
		}

	})
	log.Fatal(app.Listen(":7070"))
}
