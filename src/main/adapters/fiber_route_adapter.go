package adapters

import (
	protocols "lucaswilliameufrasio/golang-fiber-api/src/presentation/protocols"

	"github.com/gofiber/fiber/v2"
)

// AdaptRoute adapt fiber interface to any controller
func AdaptRoute(controller protocols.Controller) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var body interface{}
		var userID *int
		if err := c.BodyParser(&body); err != nil {
			c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse body",
			})
		}

		if c.Locals("userID") != nil {
			userID = nil
			id, ok := c.Locals("userID").(*int)
			if ok {
				userID = id
			}
		}

		var request = protocols.HTTPRequest{
			Body:   body,
			Params: c.Params,
			UserID: userID,
		}

		var response = controller.Handler(&request)

		return c.Status(response.StatusCode).JSON(response.Data)
	}
}
