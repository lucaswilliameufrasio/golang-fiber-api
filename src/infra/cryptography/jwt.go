package infra

import (
	"lucaswilliameufrasio/golang-fiber-api/src/data/protocols"
	"time"

	"github.com/form3tech-oss/jwt-go"
)

type JwtAdapter interface {
	protocols.Encrypter
	protocols.Decrypter
}

type JwtClaims struct {
	UserID string `json:"userID"`
	jwt.StandardClaims
}

func NewJwtAdapter(secret string) JwtAdapter {
	return JwtAdapterParams{
		secret: secret,
	}
}

type JwtAdapterParams struct {
	secret string
}

func (jwta JwtAdapterParams) Encrypt(plaintext string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = plaintext
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix()

	generatedToken, err := token.SignedString([]byte(jwta.secret))

	if err != nil {
		return "", err
	}

	return generatedToken, nil
}

func (jwta JwtAdapterParams) Decrypt(ciphertext string) (string, error) {

	token, err := jwt.ParseWithClaims(ciphertext, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwta.secret), nil
	})

	if err != nil {
		return "", err
	}

	if token.Valid {
		if claims, ok := token.Claims.(*JwtClaims); ok {
			return claims.UserID, nil
		}
	}

	return "", err
}
