package ucs

type AuthenticationParams struct {
	Email    string
	Password string
}

type AuthenticationResult struct {
	Token string
	User  struct {
		Email string
	}
}

type Authentication interface {
	Auth(p AuthenticationParams) (*AuthenticationResult, error)
}
