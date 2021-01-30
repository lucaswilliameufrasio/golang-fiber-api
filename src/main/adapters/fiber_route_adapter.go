package adapters

import (
	protocols "lucaswilliameufrasio/golang-fiber-api/src/presentation/protocols"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// AdaptRoute adapt fiber interface to any controller
func AdaptRoute(controller protocols.Controller) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var body interface{}
		c.BodyParser(&body)

		userInfo := getUserFromContext(c)

		var request = protocols.HTTPRequest{
			Body:   body,
			Params: c.Params,
			User:   userInfo,
		}

		var response = controller.Handler(&request)

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