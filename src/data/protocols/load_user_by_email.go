package protocols

type LoadUserByIDRepositoryResult struct {
	Email    string
	Password string
	Role     string
}

type LoadUserByEmailRepository interface {
	LoadByEmail(email string) (user *LoadUserByIDRepositoryResult, err error)
}