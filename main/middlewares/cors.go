package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Cors is a middleware to enable Cross-Origin Resource Sharing
func Cors(app *fiber.App) {
	app.Use(cors.New())
}
