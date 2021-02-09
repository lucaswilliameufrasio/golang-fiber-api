package fvalidators

import (
	infra "lucaswilliameufrasio/golang-fiber-api/src/infra/cryptography"
	"lucaswilliameufrasio/golang-fiber-api/src/main/config/environment"
	valiprotocols "lucaswilliameufrasio/golang-fiber-api/src/validation/protocols"
	"lucaswilliameufrasio/golang-fiber-api/src/validation/validators"
)

func MakeTokenValidation() valiprotocols.TokenValidator {
	jwtAdapter := infra.NewJwtAdapter(environment.JWT_SECRET)
	return validators.NewTokenValidation(jwtAdapter)
}
