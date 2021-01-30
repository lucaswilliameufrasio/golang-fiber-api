package fctsucs

import (
	aucs "lucaswilliameufrasio/golang-fiber-api/src/data/usecases"
	ucs "lucaswilliameufrasio/golang-fiber-api/src/domain/usecases"
	infra "lucaswilliameufrasio/golang-fiber-api/src/infra/cryptography"
	"lucaswilliameufrasio/golang-fiber-api/src/infra/repositories"
	"lucaswilliameufrasio/golang-fiber-api/src/main/config/environment"
)

func MakeDbAuthentication() ucs.Authentication {
	jwtAdapter := infra.NewJwtAdapter(environment.JWT_SECRET)
	loadUserByIDRepository := repositories.NewFakeLoadUserByIDRepository()
	hashComparer := infra.NewArgonAdapter()
	return aucs.NewDbAuthentication(jwtAdapter, loadUserByIDRepository, hashComparer)
}
