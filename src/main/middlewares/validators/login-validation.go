package validators

import (
	"errors"
	"fmt"

	presenterrors "lucaswilliameufrasio/golang-fiber-api/src/presentation/errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type LoginParams struct {
	Email    string `validate:"required,email,omitempty"`
	Password string `validate:"required,omitempty"`
}

func LoginValidation(c *fiber.Ctx) error {
	body := new(LoginParams)
	if err := c.BodyParser(&body); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse body",
		})
	}

	validate := validator.New()

	params := &LoginParams{
		Email:    body.Email,
		Password: body.Password,
	}

	err := validate.Struct(params)

	if err != nil {
		fmt.Println(err)
		for _, errorValue := range err.(validator.ValidationErrors) {
			paramKey := errors.New(errorValue.Field())
			return c.Status(400).JSON(fiber.Map{
				"error": presenterrors.MissingParamError(paramKey.Error()).Error(),
			})
		}

	}

	return c.Next()
}
