package fmdlwrs

import (
	fvalidators "lucaswilliameufrasio/golang-fiber-api/src/main/factories/validators"
	"lucaswilliameufrasio/golang-fiber-api/src/presentation/middlewares"
	presprotcls "lucaswilliameufrasio/golang-fiber-api/src/presentation/protocols"
)

func MakeAuthMiddleware() presprotcls.Middleware {
	return middlewares.NewAuthMiddleware(fvalidators.MakeTokenValidation())
}
