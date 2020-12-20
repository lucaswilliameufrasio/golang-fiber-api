package adapters

import (
	protocols "lucaswilliameufrasio/golang-fiber-api/presentation/protocols"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// AdaptRouteResult specifies what will be returned by the function
type AdaptRouteResult func(c *fiber.Ctx) error

// AdaptRoute adapt fiber interface to any controller
func AdaptRoute(controller protocols.Controller) AdaptRouteResult {
	return func(c *fiber.Ctx) error {
		var body map[string]string
		c.BodyParser(&body)

		userInfo := getUserFromContext(c)

		var request = protocols.HTTPRequest{
			Body:   body,
			Params: c.Params,
			User:   userInfo,
		}

		var response = controller(&request)

		return c.Status(response.StatusCode).JSON(response.Data)
	}
}

func getUserFromContext(c *fiber.Ctx) map[string]string {
	var id string
	var name string
	var role string

	var requestScopeVariable = c.Locals("user")
	if requestScopeVariable != nil {
		var userFromContext = c.Locals("user").(*jwt.Token)
		var claims = userFromContext.Claims.(jwt.MapClaims)

		name = claims["name"].(string)
		role = claims["role"].(string)
	}

	return map[string]string{
		"id":   id,
		"name": name,
		"role": role,
	}
}
