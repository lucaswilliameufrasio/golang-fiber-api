package config

import (
	"lucaswilliameufrasio/golang-fiber-api/src/main/middlewares"

	"github.com/gofiber/fiber/v2"
)

// SetupMiddlewares will contain all app middlewares setup
func SetupMiddlewares(app *fiber.App) {
	middlewares.Cors(app)
	middlewares.LimitRequest(app)
	middlewares.Recover(app)
}
