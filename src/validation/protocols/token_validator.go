package protocols

type TokenValidator interface {
	Validate(token string) (*string, error)
}
