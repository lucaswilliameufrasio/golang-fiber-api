package adapters

import (
	premiddlewares "lucaswilliameufrasio/golang-fiber-api/src/presentation/middlewares"
	protocols "lucaswilliameufrasio/golang-fiber-api/src/presentation/protocols"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// AdaptMiddleware adapt fiber interface to any middleware
func AdaptMiddleware(middleware protocols.Middleware) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var request = protocols.HTTPRequest{
			Token: GetTokenFromHeader(c),
		}

		var response = middleware.Handler(&request)

		if response.StatusCode == 200 {
			mapResponse, ok := response.Data.(premiddlewares.AuthMiddlewareResult)
			if !ok {
				c.Locals("userID", nil)
			} else {
				c.Locals("userID", mapResponse.ID)
			}
		} else {
			return c.Status(response.StatusCode).JSON(response.Data)
		}
		return c.Next()
	}
}

func GetTokenFromHeader(c *fiber.Ctx) string {
	tokenOnHeader := strings.Split(string(c.Request().Header.Peek("Authorization")), " ")

	if len(tokenOnHeader) == 2 {
		return tokenOnHeader[1]
	}

	return ""
}
