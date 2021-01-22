package routes

import (
	"lucaswilliameufrasio/golang-fiber-api/src/presentation/controllers"

	"github.com/gofiber/fiber/v2"
)

// LoginRoutes is a function to setup login routes
func LoginRoutes(router fiber.Router) {
	router.Post("/login", controllers.LoginController)
}
