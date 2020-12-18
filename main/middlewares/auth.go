package middlewares

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// JwtSecret will have the information of the secret needed to create and verify jwt tokens
var JwtSecret = os.Getenv("JWT_SECRET")

func handleAuthError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}

// AuthenticationRequired is a middleware to verify the token validation
func AuthenticationRequired() func(c *fiber.Ctx) error {
	var jwtConfig = jwtware.Config{
		ErrorHandler: handleAuthError,
		SigningKey:   []byte(JwtSecret),
	}
	return jwtware.New(jwtConfig)
}

// AuthorizationMiddleware is a middleware to ensure user has a role setted
func AuthorizationMiddleware(c *fiber.Ctx) error {
	role := c.Locals("role").(*jwt.Token)

	if role == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	fmt.Print(role)

	return c.Next()
}
