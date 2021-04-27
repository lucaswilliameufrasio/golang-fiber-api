package middlewares

import (
	presentationhelpers "lucaswilliameufrasio/golang-fiber-api/src/presentation/helpers"
	protocols "lucaswilliameufrasio/golang-fiber-api/src/presentation/protocols"
	valiprotocols "lucaswilliameufrasio/golang-fiber-api/src/validation/protocols"
	"strconv"
)

type AuthMiddlewareResult struct {
	ID int `json:"id"`
}

func NewAuthMiddleware(validator valiprotocols.TokenValidator) protocols.Middleware {
	return AuthMiddleware{
		validator,
	}
}

type AuthMiddleware struct {
	valiprotocols.TokenValidator
}

func (a AuthMiddleware) Handler(request *protocols.HTTPRequest) protocols.HTTPResponse {
	id, err := a.Validate(request.Token)

	if err != nil {
		return presentationhelpers.Unauthorized()
	}

	userID, err := strconv.Atoi(*id)

	if err != nil {
		return presentationhelpers.Unauthorized()
	}

	return presentationhelpers.OK(AuthMiddlewareResult{
		ID: userID,
	})
}
