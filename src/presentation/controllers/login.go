package controllers

import (
	ucs "lucaswilliameufrasio/golang-fiber-api/src/domain/usecases"
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
	convertedData := data.(map[string]interface{})
	email := convertedData["email"].(string)
	password := convertedData["password"].(string)
	params := LoginControllerParams{
		Email:    email,
		Password: password,
	}
	return params
}

// Handler is a controller to execute login process
func (sts LoginController) Handler(request *protocols.HTTPRequest) protocols.HTTPResponse {
	params := castParams(request.Body)

	result, err := sts.Authentication.Auth(ucs.AuthenticationParams(params))

	if err != nil || result == nil {
		return protocols.HTTPResponse{
			StatusCode: 401,
			Data: map[string]interface{}{
				"error": "Bad Credentials",
			},
		}
	}

	return protocols.HTTPResponse{
		StatusCode: 200,
		Data: LoginControllerResult{
			Token: result.Token,
			Email: result.User.Email,
		},
	}
}
