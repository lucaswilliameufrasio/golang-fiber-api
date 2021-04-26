package controllers

import (
	ucs "lucaswilliameufrasio/golang-fiber-api/src/domain/usecases"
	presenterrors "lucaswilliameufrasio/golang-fiber-api/src/presentation/errors"
	presenthelpers "lucaswilliameufrasio/golang-fiber-api/src/presentation/helpers"
	protocols "lucaswilliameufrasio/golang-fiber-api/src/presentation/protocols"
)

type LoginControllerParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

func castParams(data interface{}) LoginControllerParams {
	dataToMap := data.(map[string]interface{})

	if dataToMap["email"] == nil || dataToMap["password"] == nil {
		return LoginControllerParams{
			Email:    "",
			Password: "",
		}
	}

	email := dataToMap["email"].(string)
	password := dataToMap["password"].(string)
	params := LoginControllerParams{
		Email:    email,
		Password: password,
	}
	return params
}

// Handler is a controller to execute login process
func (sts LoginController) Handler(request *protocols.HTTPRequest) protocols.HTTPResponse {
	if request.Body == nil {
		return presenthelpers.BadRequest(presenterrors.MissingParamError("email"))
	}
	params := castParams(request.Body)

	if params.Email == "" {
		return presenthelpers.BadRequest(presenterrors.MissingParamError("email"))
	}

	if params.Password == "" {
		return presenthelpers.BadRequest(presenterrors.MissingParamError("password"))
	}

	result, err := sts.Authentication.Auth(ucs.AuthenticationParams(params))

	if err != nil || result == nil {
		return presenthelpers.Unauthorized()
	}

	return protocols.HTTPResponse{
		StatusCode: 200,
		Data: LoginControllerResult{
			Token: result.Token,
			Email: result.User.Email,
		},
	}
}
