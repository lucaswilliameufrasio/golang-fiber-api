package controllers

import "github.com/gofiber/fiber/v2"

// ReplyGreetingsController is a controller to handle request and respond with a great message
func ReplyGreetingsController(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
