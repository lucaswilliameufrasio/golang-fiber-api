package routes

import (
	"lucaswilliameufrasio/golang-fiber-api/src/main/adapters"
	"lucaswilliameufrasio/golang-fiber-api/src/presentation/controllers"

	"github.com/gofiber/fiber/v2"
)

// StatusRoutes setup
func StatusRoutes(router fiber.Router) {
	router.Get("/", adapters.AdaptRoute(controllers.StatusController))
}
