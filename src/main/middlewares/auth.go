package middlewares

import (
	"errors"
	"fmt"
	"log"
	"lucaswilliameufrasio/golang-fiber-api/src/main/config/environment"
	"strings"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func handleAuthError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}

func GetTokenFromHeader(c *fiber.Ctx) (string, error) {
	tokenOnHeader := strings.Split(string(c.Request().Header.Peek("Authorization")), " ")

	if len(tokenOnHeader) == 2 {
		return tokenOnHeader[1], nil
	}

	return "", errors.New("Can't find any token")
}

func AuthenticationMiddleware() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		tokenFromHeader, err := GetTokenFromHeader(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		type JWTClaims struct {
			UserID string `json:"userID"`
			jwt.StandardClaims
		}

		token, err := jwt.ParseWithClaims(tokenFromHeader, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(environment.JWT_SECRET), nil
		})

		if token.Valid {
			if claims, ok := token.Claims.(*JWTClaims); ok {
				fmt.Printf("%v\n", claims.UserID)
				return c.Next()
			} else {
				log.Fatal(err)
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Unauthorized",
				})
			}
		}

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
}

// AuthenticationRequired is a middleware to verify the token validation
func AuthenticationRequired() func(c *fiber.Ctx) error {
	var jwtConfig = jwtware.Config{
		ErrorHandler: handleAuthError,
		SigningKey:   []byte(environment.JWT_SECRET),
	}
	return jwtware.New(jwtConfig)
}

// AuthorizationMiddleware is a middleware to ensure user has a role setted
// func AuthorizationMiddleware() func(*fiber.Ctx) error {
// 	return func(c *fiber.Ctx) error {
// 		var role string

// 		var requestScopeVariable = c.Locals("user")
// 		if requestScopeVariable == nil {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": "Unauthorized",
// 			})
// 		}
// 		var userFromContext = c.Locals("user").(*jwt.Token)
// 		var claims = userFromContext.Claims.(jwt.MapClaims)

// 		role = claims["role"].(string)

// 		if role == "guest" {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": "Unauthorized",
// 			})
// 		}

// 		return c.Next()
// 	}
// }
