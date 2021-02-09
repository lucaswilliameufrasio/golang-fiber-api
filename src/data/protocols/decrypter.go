package protocols

type Decrypter interface {
	Decrypt(ciphertext string) (string, error)
}
