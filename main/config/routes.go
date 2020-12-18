package config

import (
	"lucaswilliameufrasio/golang-fiber-api/main/events"
	"lucaswilliameufrasio/golang-fiber-api/main/middlewares"
	"lucaswilliameufrasio/golang-fiber-api/main/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

// SetupRoutes is a factory functions to setup routes
func SetupRoutes(app *fiber.App) fiber.Router {
	var router = app.Group("/api")

	router.Get("/", routes.Greetings)
	router.Get("/ws", websocket.New(events.WebsocketHandler))

	setupRoutesV1(router)
	setupRoutesV2(router)

	return router
}

// SetupRoutesV1 is a function to export app version 1 routes
func setupRoutesV1(router fiber.Router) fiber.Router {
	v1 := router.Group("/v1", middlewares.SimpleMiddleware)

	// GET /john/75
	v1.Get("/profile/:name/:age/:gender?", routes.UserNameAndAge)

	return v1
}

// SetupRoutesV2 is a function to export app version 2 routes
func setupRoutesV2(router fiber.Router) fiber.Router {
	v2 := router.Group("/v2")

	v2.Use(middlewares.LimitRequest)

	v2.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	v2.Post("/login", routes.Login)

	v2.Get("/profile", middlewares.AuthenticationRequired(), routes.Profile)

	return v2
}
