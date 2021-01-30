package protocols

type HashComparer interface {
	Compare(plaintext string, digest string) (bool, error)
}
