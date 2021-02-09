package middlewares

import (
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
		return protocols.HTTPResponse{
			StatusCode: 403,
			Data: map[string]interface{}{
				"error": "Access Denied",
			},
		}
	}

	userID, err := strconv.Atoi(*id)

	if err != nil {
		return protocols.HTTPResponse{
			StatusCode: 403,
			Data: map[string]interface{}{
				"error": "Access Denied",
			},
		}
	}

	return protocols.HTTPResponse{
		StatusCode: 200,
		Data: AuthMiddlewareResult{
			ID: userID,
		},
	}
}
