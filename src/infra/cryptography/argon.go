package infra

import (
	"lucaswilliameufrasio/golang-fiber-api/src/data/protocols"

	"github.com/matthewhartstonge/argon2"
)

func NewArgonAdapter() protocols.HashComparer {
	return ArgonAdapter{}
}

type ArgonAdapter struct {
}

func (argona ArgonAdapter) Compare(plaintext string, digest string) (bool, error) {
	ok, err := argon2.VerifyEncoded([]byte(plaintext), []byte(digest))
	if err != nil {
		return false, err
	}

	if ok {
		return true, nil
	}

	return false, nil
}
