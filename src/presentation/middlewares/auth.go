package middlewares

import (
	ucs "lucaswilliameufrasio/golang-fiber-api/src/domain/usecases"
	protocols "lucaswilliameufrasio/golang-fiber-api/src/presentation/protocols"
)

type AuthMiddlewareResult struct {
	ID int `json:"id"`
}

func NewAuthMiddleware(load ucs.LoadUserIDByToken) protocols.Middleware {
	return AuthMiddleware{
		load,
	}
}

type AuthMiddleware struct {
	ucs.LoadUserIDByToken
}

func (a AuthMiddleware) Handler(request *protocols.HTTPRequest) protocols.HTTPResponse {
	userId, err := a.LoadUserIDByToken.Load(request.Token)

	if err != nil {
		return protocols.HTTPResponse{
			StatusCode: 403,
			Data: map[string]interface{}{
				"error": "AccessDeniedError",
			},
		}
	}

	return protocols.HTTPResponse{
		StatusCode: 200,
		Data: AuthMiddlewareResult{
			ID: userId.ID,
		},
	}
}
