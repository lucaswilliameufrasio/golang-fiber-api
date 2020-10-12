package main

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/gofiber/websocket/v2"
)

const jwtSecret = "golang"

func middleware(c *fiber.Ctx) error {
	fmt.Println("May the knowledge be with you")
	return c.Next()
}

func wsHandler(c *websocket.Conn) {
	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", msg)
		err = c.WriteMessage(mt, msg)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func server() *fiber.App {
	app := fiber.New()

	app.Use(cors.New())
	api := app.Group("/api")
	v1 := api.Group("/v1", middleware)
	v2 := api.Group("/v2")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	v1.Static("/static", "./public")

	// GET /john/75
	v1.Get("/profile/:name/:age/:gender?", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))
		return c.JSON(fiber.Map{
			"info":   msg,
			"active": true,
		}) // => info: 👴 john is 75 years old
	})

	v2.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:      20,
		Duration: 30 * time.Second,
		Key: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"error": "Slow down your fingers, mate.",
			})
		},
	}))

	v2.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	v2.Post("/login", login)

	v2.Get("/profile", authRequired(), profile)

	app.Get("/ws", websocket.New(wsHandler))

	return app
}

func main() {
	app := server()
	// app.Use(logger.New())
	log.Fatal(app.Listen(":7777"))
}

func login(c *fiber.Ctx) error {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body request
	err := c.BodyParser(&body)

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse json",
		})
	}

	if body.Email != "john@example.com" || body.Password != "doe" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Bad Credentials",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "John Doe"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix()

	generatedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": generatedToken,
		"user": struct {
			ID    int    `json:"id"`
			Email string `json:"email"`
		}{
			ID:    1,
			Email: "john@example.com",
		},
	})
}

func authRequired() func(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
		SigningKey: []byte(jwtSecret),
	})
}

func authMiddleware(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)

	if user == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}

func profile(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	return c.SendString("Welcome, " + name)
}
