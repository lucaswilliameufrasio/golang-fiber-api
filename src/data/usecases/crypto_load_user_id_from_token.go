package aucs

import (
	"lucaswilliameufrasio/golang-fiber-api/src/data/protocols"
	ucs "lucaswilliameufrasio/golang-fiber-api/src/domain/usecases"
	"strconv"
)

func NewCryptoLoadUserByToken(dec protocols.Decrypter) ucs.LoadUserIDByToken {
	return CryptoLoadUserByTokenResult{
		dec,
	}
}

type CryptoLoadUserByTokenResult struct {
	protocols.Decrypter
}

func (a CryptoLoadUserByTokenResult) Load(Token string) (*ucs.LoadUserIDByTokenResult, error) {
	id, err := a.Decrypter.Decrypt(Token)

	if err != nil {
		return nil, err
	}

	if id != "" {
		userId, err := strconv.Atoi(id)
		if err != nil {
			return nil, err
		}
		return &ucs.LoadUserIDByTokenResult{
			ID: userId,
		}, nil
	}

	return nil, nil
}
