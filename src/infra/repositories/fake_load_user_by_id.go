package repositories

import (
	"lucaswilliameufrasio/golang-fiber-api/src/data/protocols"
	datasources "lucaswilliameufrasio/golang-fiber-api/src/infra/datasources"
)

func NewFakeLoadUserByIDRepository() protocols.LoadUserByEmailRepository {
	return fakeLoadUserByIDRepository{}
}

type fakeLoadUserByIDRepository struct{}

func (fakeLoadUserByIDRepository) LoadByEmail(email string) (*protocols.LoadUserByIDRepositoryResult, error) {
	users := datasources.MakeUsersDataSource()
	for _, user := range users {
		if user.Email == email {
			return &protocols.LoadUserByIDRepositoryResult{
				ID:       user.ID,
				Email:    user.Email,
				Password: user.Password,
				Role:     user.Role,
			}, nil
		}
	}
	return nil, nil
}
