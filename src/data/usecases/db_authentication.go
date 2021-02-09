package aucs

import (
	"fmt"
	"lucaswilliameufrasio/golang-fiber-api/src/data/protocols"
	ucs "lucaswilliameufrasio/golang-fiber-api/src/domain/usecases"
)

func NewDbAuthentication(enc protocols.Encrypter, repo protocols.LoadUserByEmailRepository, comparer protocols.HashComparer) ucs.Authentication {
	return dbAuthentication{
		enc,
		repo,
		comparer,
	}
}

type dbAuthentication struct {
	protocols.Encrypter
	protocols.LoadUserByEmailRepository
	protocols.HashComparer
}

func (a dbAuthentication) Auth(p ucs.AuthenticationParams) (*ucs.AuthenticationResult, error) {
	account, err := a.LoadUserByEmailRepository.LoadByEmail(p.Email)

	if err != nil {
		return nil, err
	}

	if account != nil {
		var iseq bool
		if iseq, err = a.HashComparer.Compare(p.Password, account.Password); err != nil {
			return nil, err
		}

		if iseq == true {
			userID := fmt.Sprintf("%v", account.ID)
			generatedToken, err := a.Encrypt(userID)

			if err != nil {
				return nil, err
			}

			return &ucs.AuthenticationResult{
				Token: generatedToken,
				User: struct{ Email string }{
					Email: p.Email,
				},
			}, nil
		}

	}

	return nil, nil
}
