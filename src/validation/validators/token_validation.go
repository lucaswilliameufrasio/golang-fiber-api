package validators

import (
	"lucaswilliameufrasio/golang-fiber-api/src/data/protocols"
	valiprotocols "lucaswilliameufrasio/golang-fiber-api/src/validation/protocols"
)

func NewTokenValidation(dec protocols.Decrypter) valiprotocols.TokenValidator {
	return TokenValidationResult{
		dec,
	}
}

type TokenValidationResult struct {
	protocols.Decrypter
}

func (a TokenValidationResult) Validate(token string) (*string, error) {
	id, err := a.Decrypter.Decrypt(token)
	if err != nil {
		return nil, err
	}
	if id != "" {
		return &id, nil
	}

	return nil, nil
}
