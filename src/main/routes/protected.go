package routes

import (
	"lucaswilliameufrasio/golang-fiber-api/src/main/adapters"
	fctrls "lucaswilliameufrasio/golang-fiber-api/src/main/factories/controllers"
	fmdlwrs "lucaswilliameufrasio/golang-fiber-api/src/main/factories/middlewares"

	"github.com/gofiber/fiber/v2"
)

// Protected setup
func Protected(router fiber.Router) {
	router.Get("/protected", adapters.AdaptMiddleware(fmdlwrs.MakeAuthMiddleware()), adapters.AdaptRoute(fctrls.MakeProtectedControler()))
}
