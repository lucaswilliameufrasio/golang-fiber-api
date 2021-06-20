package controllers

import (
	"errors"
	ucs "lucaswilliameufrasio/golang-fiber-api/src/domain/usecases"
	presenterrors "lucaswilliameufrasio/golang-fiber-api/src/presentation/errors"
	presenthelpers "lucaswilliameufrasio/golang-fiber-api/src/presentation/helpers"
	protocols "lucaswilliameufrasio/golang-fiber-api/src/presentation/protocols"

	"github.com/go-playground/validator/v10"
)

type LoginControllerParams struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginControllerResult struct {
	Token string `json:"token"`
	Email string `json:"email"`
}

func NewLoginController(auth ucs.Authentication) protocols.Controller {
	return LoginController{
		auth,
	}
}

type LoginController struct {
	ucs.Authentication
}

func castAndValidateParams(data interface{}) (*LoginControllerParams, error) {
	dataToMap := data.(map[string]interface{})

	validate := validator.New()

	email := dataToMap["email"].(string)
	password := dataToMap["password"].(string)

	params := &LoginControllerParams{
		Email:    string(email),
		Password: password,
	}

	err := validate.Struct(params)

	if err != nil {
		var paramKey error
		for _, errorValue := range err.(validator.ValidationErrors) {
			paramKey = errors.New(errorValue.Tag())
			return nil, paramKey
		}

	}

	return params, nil
}

// Handler is a controller to execute login process
func (sts LoginController) Handler(request *protocols.HTTPRequest) protocols.HTTPResponse {
	if request.Body == nil {
		return presenthelpers.BadRequest(presenterrors.MissingParamError("email"))
	}
	params, err := castAndValidateParams(request.Body)

	if err != nil {
		return presenthelpers.BadRequest(presenterrors.MissingParamError(err.Error()))
	}

	result, err := sts.Authentication.Auth(ucs.AuthenticationParams(*params))

	if err != nil || result == nil {
		return presenthelpers.Unauthorized()
	}

	return presenthelpers.OK(LoginControllerResult{
		Token: result.Token,
		Email: result.User.Email,
	})
}
