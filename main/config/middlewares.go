package config

import (
	"lucaswilliameufrasio/golang-fiber-api/main/middlewares"

	"github.com/gofiber/fiber/v2"
)

// SetupMiddlewares will contain all app middlewares setup
func SetupMiddlewares(app *fiber.App) {
	middlewares.Cors(app)
}
