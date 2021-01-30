package routes

import (
	"lucaswilliameufrasio/golang-fiber-api/src/main/adapters"
	fctrls "lucaswilliameufrasio/golang-fiber-api/src/main/factories/controllers"
	"lucaswilliameufrasio/golang-fiber-api/src/main/middlewares"

	"github.com/gofiber/fiber/v2"
)

// Protected setup
func Protected(router fiber.Router) {
	router.Get("/protected", middlewares.AuthenticationRequired(), middlewares.AuthorizationMiddleware(), adapters.AdaptRoute(fctrls.MakeProtectedControler()))
}
