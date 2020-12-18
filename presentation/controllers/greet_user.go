package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// GreetUserController is intended to return greetings message
func GreetUserController(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	return c.SendString("Welcome, " + name)
}
