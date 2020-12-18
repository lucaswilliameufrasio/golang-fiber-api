package routes

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Profile is a route to pass data to controller and return a greetings message
func Profile(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	return c.SendString("Welcome, " + name)
}

// UserNameAndAge to get user name and age
func UserNameAndAge(c *fiber.Ctx) error {
	msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))
	return c.JSON(fiber.Map{
		"info":   msg,
		"active": true,
	}) // => info: 👴 john is 75 years old
}
