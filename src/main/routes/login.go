package routes

import (
	"lucaswilliameufrasio/golang-fiber-api/src/main/adapters"
	fctrls "lucaswilliameufrasio/golang-fiber-api/src/main/factories/controllers"

	"github.com/gofiber/fiber/v2"
)

// LoginRoutes is a function to setup login routes
func LoginRoutes(router fiber.Router) {
	router.Post("/login", adapters.AdaptRoute(fctrls.MakeLoginController()))
}
