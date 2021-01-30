package infra

import (
	"lucaswilliameufrasio/golang-fiber-api/src/data/protocols"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func NewJwtAdapter(secret string) protocols.Encrypter {
	return JwtAdapter{
		secret: secret,
	}
}

type JwtAdapter struct {
	secret string
}

func (jwta JwtAdapter) Encrypt(plaintext string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["role"] = plaintext
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix()

	generatedToken, err := token.SignedString([]byte(jwta.secret))

	if err != nil {
		return "", err
	}

	return generatedToken, nil
}
