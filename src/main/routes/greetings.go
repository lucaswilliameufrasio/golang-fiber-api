package routes

import (
	"lucaswilliameufrasio/golang-fiber-api/src/main/adapters"
	"lucaswilliameufrasio/golang-fiber-api/src/presentation/controllers"

	"github.com/gofiber/fiber/v2"
)

// GreetingsRoutes setup
func GreetingsRoutes(router fiber.Router) {
	router.Get("/", adapters.AdaptRoute(controllers.ReplyGreetingsController))
}
