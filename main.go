package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetRequests(c *fiber.Ctx) error {

	getString := c.OriginalURL()

	UpperCaseString := strings.ToUpper(getString)

	return c.SendString(UpperCaseString)
}

func NoEndpoint(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
}

func main() {
	app := fiber.New()

	app.Get("/", NoEndpoint)

	app.Get("/:text?", GetRequests)

	log.Fatal(app.Listen(":7070"))
}
