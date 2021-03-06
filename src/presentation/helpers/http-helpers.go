package presentationhelpers

import (
	errors "lucaswilliameufrasio/golang-fiber-api/src/presentation/errors"
	protocols "lucaswilliameufrasio/golang-fiber-api/src/presentation/protocols"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func Unauthorized() protocols.HTTPResponse {
	return protocols.HTTPResponse{
		StatusCode: 401,
		Data: ErrorResponse{
			Error: errors.UnauthorizedError().Error(),
		},
	}
}

func OK(data interface{}) protocols.HTTPResponse {
	return protocols.HTTPResponse{
		StatusCode: 200,
		Data:       data,
	}
}
