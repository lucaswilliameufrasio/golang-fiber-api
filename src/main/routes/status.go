package routes

import (
	"lucaswilliameufrasio/golang-fiber-api/src/main/adapters"
	fctrls "lucaswilliameufrasio/golang-fiber-api/src/main/factories/controllers"

	"github.com/gofiber/fiber/v2"
)

// StatusRoutes setup
func StatusRoutes(router fiber.Router) {
	router.Get("/", adapters.AdaptRoute(fctrls.MakeStatusController()))
}
