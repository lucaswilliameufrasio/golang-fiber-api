// +test

package aucs_test

import (
	"errors"
	mocks "lucaswilliameufrasio/golang-fiber-api/src/data/mocks"
	aucs "lucaswilliameufrasio/golang-fiber-api/src/data/usecases"
	ucs "lucaswilliameufrasio/golang-fiber-api/src/domain/usecases"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	Faker          faker.Faker
	MockController *gomock.Controller
)

type SutTypes struct {
	sut                           ucs.Authentication
	mockEncrypter                 *mocks.MockEncrypter
	mockLoadUserByEmailRepository *mocks.MockLoadUserByEmailRepository
	mockHashComparer              *mocks.MockHashComparer
}

func SUT(t *testing.T) SutTypes {
	Faker = faker.New()
	MockController = gomock.NewController(t)
	defer MockController.Finish()
	mockEncrypter := mocks.NewMockEncrypter(MockController)
	mockLoadUserByEmailRepository := mocks.NewMockLoadUserByEmailRepository(MockController)
	mockHashComparer := mocks.NewMockHashComparer(MockController)

	mockHashComparer.EXPECT().Compare(gomock.Any().String(), gomock.Any().String()).AnyTimes()
	mockEncrypter.EXPECT().Encrypt(gomock.Any().String()).AnyTimes()
	mockLoadUserByEmailRepository.EXPECT().LoadByEmail(gomock.Any().String()).AnyTimes()

	sut := aucs.NewDbAuthentication(mockEncrypter, mockLoadUserByEmailRepository, mockHashComparer)

	return SutTypes{
		sut:                           sut,
		mockEncrypter:                 mockEncrypter,
		mockLoadUserByEmailRepository: mockLoadUserByEmailRepository,
		mockHashComparer:              mockHashComparer,
	}
}

func TestCallLoadUserByEmailRepositoryCorrectly(t *testing.T) {
	instances := SUT(t)

	authenticationParams := ucs.AuthenticationParams{
		Email:    Faker.Internet().Email(),
		Password: Faker.Internet().Password(),
	}

	instances.mockLoadUserByEmailRepository.EXPECT().LoadByEmail(gomock.Eq(authenticationParams.Email)).AnyTimes()

	sut := instances.sut

	_, err := sut.Auth(authenticationParams)

	assert.Equal(t, err, nil)
}

func TestThrowIfLoadUserByEmailRepositoryThrows(t *testing.T) {
	instances := SUT(t)

	authenticationParams := ucs.AuthenticationParams{
		Email:    Faker.Internet().Email(),
		Password: Faker.Internet().Password(),
	}

	instances.mockLoadUserByEmailRepository.EXPECT().LoadByEmail(gomock.Eq(authenticationParams.Email)).Return(nil, errors.New("Generic"))

	sut := instances.sut

	result, err := sut.Auth(authenticationParams)

	assert.Nil(t, result)
	assert.Equal(t, err, errors.New("Generic"))
}
