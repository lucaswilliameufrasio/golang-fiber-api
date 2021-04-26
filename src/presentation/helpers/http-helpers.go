package presentationhelpers

import (
	errors "lucaswilliameufrasio/golang-fiber-api/src/presentation/errors"
	protocols "lucaswilliameufrasio/golang-fiber-api/src/presentation/protocols"
)

func BadRequest(errorParam error) protocols.HTTPResponse {
	return protocols.HTTPResponse{
		StatusCode: 400,
		Data:       errorParam.Error(),
	}
}

func Unauthorized() protocols.HTTPResponse {
	return protocols.HTTPResponse{
		StatusCode: 401,
		Data:       errors.UnauthorizedError(),
	}
}
