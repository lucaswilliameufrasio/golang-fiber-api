package config

import (
	"lucaswilliameufrasio/golang-fiber-api/src/main/events"
	"lucaswilliameufrasio/golang-fiber-api/src/main/middlewares"
	"lucaswilliameufrasio/golang-fiber-api/src/main/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

// SetupRoutes is a factory functions to setup routes
func SetupRoutes(app *fiber.App) fiber.Router {
	var router = app.Group("/api")

	routes.GreetingsRoutes(router)

	setupRoutesV1(router)
	setupRoutesV2(router)

	router.Get("/ws", websocket.New(events.WebsocketHandler))

	return router
}

// SetupRoutesV1 is a function to export app version 1 routes
func setupRoutesV1(router fiber.Router) fiber.Router {
	v1 := router.Group("/v1", middlewares.SimpleMiddleware)

	routes.UserRoutes(v1)

	return v1
}

// SetupRoutesV2 is a function to export app version 2 routes
func setupRoutesV2(router fiber.Router) fiber.Router {
	v2 := router.Group("/v2")

	v2.Use(middlewares.LimitRequest)

	v2.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	routes.LoginRoutes(v2)

	return v2
}
