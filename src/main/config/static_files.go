package config

import "github.com/gofiber/fiber/v2"

// SetupStaticFiles will set a route to access files on public folder
func SetupStaticFiles(app *fiber.App) {
	app.Static("/static", "./public")
}
