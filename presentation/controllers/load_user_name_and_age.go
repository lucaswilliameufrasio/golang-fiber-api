package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// LoadUserNameAndAgeController to get user name and age
func LoadUserNameAndAgeController(c *fiber.Ctx) error {
	msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))

	return c.JSON(fiber.Map{
		"info":   msg,
		"active": true,
	}) // => info: 👴 john is 75 years old
}
