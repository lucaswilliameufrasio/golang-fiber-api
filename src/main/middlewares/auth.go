package middlewares

import (
	"lucaswilliameufrasio/golang-fiber-api/src/main/config/environment"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func handleAuthError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}

// AuthenticationRequired is a middleware to verify the token validation
func AuthenticationRequired() func(c *fiber.Ctx) error {
	var jwtConfig = jwtware.Config{
		ErrorHandler: handleAuthError,
		SigningKey:   []byte(environment.JWT_SECRET),
	}
	return jwtware.New(jwtConfig)
}

// AuthorizationMiddleware is a middleware to ensure user has a role setted
func AuthorizationMiddleware() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var role string

		var requestScopeVariable = c.Locals("user")
		if requestScopeVariable == nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		var userFromContext = c.Locals("user").(*jwt.Token)
		var claims = userFromContext.Claims.(jwt.MapClaims)

		role = claims["role"].(string)

		if role == "guest" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		return c.Next()
	}
}
