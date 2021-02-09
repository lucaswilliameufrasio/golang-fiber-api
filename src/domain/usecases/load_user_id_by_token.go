package ucs

type LoadUserIDByTokenResult struct {
	ID int
}

type LoadUserIDByToken interface {
	Load(Token string) (*LoadUserIDByTokenResult, error)
}
