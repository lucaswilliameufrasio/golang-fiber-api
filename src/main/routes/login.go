package routes

import (
	"lucaswilliameufrasio/golang-fiber-api/src/main/adapters"
	fctrls "lucaswilliameufrasio/golang-fiber-api/src/main/factories/controllers"
	"lucaswilliameufrasio/golang-fiber-api/src/main/middlewares/validators"

	"github.com/gofiber/fiber/v2"
)

// LoginRoutes is a function to setup login routes
func LoginRoutes(router fiber.Router) {
	router.Post("/login", validators.LoginValidation, adapters.AdaptRoute(fctrls.MakeLoginController()))
}
