package protocols

type LoadUserByIDRepositoryResult struct {
	ID       int
	Email    string
	Password string
	Role     string
}

type LoadUserByEmailRepository interface {
	LoadByEmail(email string) (*LoadUserByIDRepositoryResult, error)
}
