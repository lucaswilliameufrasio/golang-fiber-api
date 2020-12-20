package routes

import (
	"lucaswilliameufrasio/golang-fiber-api/main/adapters"
	"lucaswilliameufrasio/golang-fiber-api/presentation/controllers"

	"github.com/gofiber/fiber/v2"
)

// GreetingsRoutes setup
func GreetingsRoutes(router fiber.Router) {
	router.Get("/", adapters.AdaptRoute(controllers.ReplyGreetingsController))
}
