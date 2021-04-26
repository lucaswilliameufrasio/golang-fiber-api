package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Recover is a middleware to recover from panics anywhere in the stack chain
func Recover(app *fiber.App) {
	app.Use(recover.New())
}
