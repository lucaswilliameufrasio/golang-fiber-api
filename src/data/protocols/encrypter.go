package protocols

type Encrypter interface {
	Encrypt(plaintext string) (string, error)
}
