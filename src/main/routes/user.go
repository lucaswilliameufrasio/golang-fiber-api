package routes

import (
	"lucaswilliameufrasio/golang-fiber-api/src/main/adapters"
	"lucaswilliameufrasio/golang-fiber-api/src/main/middlewares"
	"lucaswilliameufrasio/golang-fiber-api/src/presentation/controllers"

	"github.com/gofiber/fiber/v2"
)

// UserRoutes setup
func UserRoutes(router fiber.Router) {
	// GET /john/75
	router.Get("/profile/:name/:age/:gender?", adapters.AdaptRoute(controllers.LoadUserNameAndAgeController))
	router.Get("/profile", middlewares.AuthenticationRequired(), adapters.AdaptRoute(controllers.GreetUserController))
}
