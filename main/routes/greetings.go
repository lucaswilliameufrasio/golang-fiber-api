package routes

import (
	"github.com/gofiber/fiber/v2"
)

// Greetings is a route to return a greetings message to guest client
func Greetings(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
