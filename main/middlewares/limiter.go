package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// LimitRequest is a middleware to configure a limit to client requests
func LimitRequest(c *fiber.Ctx) error {
	limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:      20,
		Duration: 30 * time.Second,
		Key: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"error": "Slow down your fingers, mate.",
			})
		},
	})

	return c.Next()
}
