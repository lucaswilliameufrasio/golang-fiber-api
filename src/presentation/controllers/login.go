package controllers

import (
	"encoding/json"
	ucs "lucaswilliameufrasio/golang-fiber-api/src/domain/usecases"
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
	jsonString, _ := json.Marshal(data)

	params := LoginControllerParams{}

	json.Unmarshal(jsonString, &params)

	return params
}

// Handler is a controller to execute login process
func (sts LoginController) Handler(request *protocols.HTTPRequest) protocols.HTTPResponse {
	params := castParams(request.Body)

	result, err := sts.Authentication.Auth(ucs.AuthenticationParams(params))

	if err != nil || result == nil {
		return presenthelpers.Unauthorized()
	}

	return presenthelpers.OK(LoginControllerResult{
		Token: result.Token,
		Email: result.User.Email,
	})
}
