package controllers

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// JwtSecret will have the information of the secret needed to create and verify jwt tokens
var JwtSecret = os.Getenv("JWT_SECRET")

// LoginController is a controller to execute login process
func LoginController(c *fiber.Ctx) error {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body request
	err := c.BodyParser(&body)

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse json",
		})
	}

	if body.Email != "john@example.com" || body.Password != "doe" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Bad Credentials",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "John Doe"
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix()

	generatedToken, err := token.SignedString([]byte(JwtSecret))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": generatedToken,
		"user": struct {
			ID    int    `json:"id"`
			Email string `json:"email"`
		}{
			ID:    1,
			Email: "john@example.com",
		},
	})
}
