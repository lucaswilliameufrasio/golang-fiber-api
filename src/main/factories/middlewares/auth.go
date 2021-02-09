package fmdlwrs

import (
	factories "lucaswilliameufrasio/golang-fiber-api/src/main/factories/usecases"
	"lucaswilliameufrasio/golang-fiber-api/src/presentation/middlewares"
	presprotcls "lucaswilliameufrasio/golang-fiber-api/src/presentation/protocols"
)

func MakeAuthMiddleware() presprotcls.Middleware {
	return middlewares.NewAuthMiddleware(factories.MakeCryptoLoadUserIDFromToken())
}
