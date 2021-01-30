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

	routes.StatusRoutes(router)

	setupRoutesV1(router)

	router.Get("/ws", websocket.New(events.WebsocketHandler))

	return router
}

// SetupRoutesV1 is a function to export app version 1 routes
func setupRoutesV1(router fiber.Router) fiber.Router {
	v1 := router.Group("/v1", middlewares.SimpleMiddleware)
	v1.Use(middlewares.LimitRequest)

	routes.LoginRoutes(v1)
	routes.Protected(v1)

	return v1
}
