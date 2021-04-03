// +test

package aucs_test

import (
	"fmt"
	mocks "lucaswilliameufrasio/golang-fiber-api/src/data/mocks"
	"lucaswilliameufrasio/golang-fiber-api/src/data/protocols"
	aucs "lucaswilliameufrasio/golang-fiber-api/src/data/usecases"
	ucs "lucaswilliameufrasio/golang-fiber-api/src/domain/usecases"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestCallLoadUserByEmailRepositoryCorrectly(t *testing.T) {
	faker := faker.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authenticationParams := ucs.AuthenticationParams{
		Email:    faker.Internet().Email(),
		Password: faker.Internet().Password(),
	}

	userId := faker.RandomDigit()
	userIdAsString := fmt.Sprintf("%v", userId)

	mockEncrypter := mocks.NewMockEncrypter(ctrl)
	mockLoadUserByEmailRepository := mocks.NewMockLoadUserByEmailRepository(ctrl)
	mockHashComparer := mocks.NewMockHashComparer(ctrl)

	mockLoadUserByEmailRepository.EXPECT().LoadByEmail(gomock.Eq(authenticationParams.Email)).Return(&protocols.LoadUserByIDRepositoryResult{
		ID:    userId,
		Email: authenticationParams.Email,
	}, nil)
	mockHashComparer.EXPECT().Compare(gomock.Eq(authenticationParams.Password), gomock.Eq("")).Return(true, nil)
	mockEncrypter.EXPECT().Encrypt(gomock.Eq(userIdAsString)).Return("", nil)
	sut := aucs.NewDbAuthentication(mockEncrypter, mockLoadUserByEmailRepository, mockHashComparer)

	_, err := sut.Auth(authenticationParams)

	assert.Equal(t, err, nil)
}
