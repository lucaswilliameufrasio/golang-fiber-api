package fctsucs

import (
	aucs "lucaswilliameufrasio/golang-fiber-api/src/data/usecases"
	ucs "lucaswilliameufrasio/golang-fiber-api/src/domain/usecases"
	infra "lucaswilliameufrasio/golang-fiber-api/src/infra/cryptography"
	"lucaswilliameufrasio/golang-fiber-api/src/main/config/environment"
)

func MakeCryptoLoadUserIDFromToken() ucs.LoadUserIDByToken {
	jwtAdapter := infra.NewJwtAdapter(environment.JWT_SECRET)
	return aucs.NewCryptoLoadUserByToken(jwtAdapter)
}
