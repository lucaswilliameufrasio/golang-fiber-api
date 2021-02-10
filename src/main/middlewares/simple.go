package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// SimpleMiddleware is a middleware only to figure out how that works on fiber
func SimpleMiddleware(c *fiber.Ctx) error {
	go fmt.Println("May the knowledge be with you")
	return c.Next()
}
