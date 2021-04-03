package config

import (
	"github.com/gofiber/fiber/v2"
)

// App is a function to start the HTTP Server
func App(config fiber.Config) *fiber.App {
	app := fiber.New(config)

	SetupStaticFiles(app)
	SetupMiddlewares(app)
	SetupRoutes(app)

	return app
}
