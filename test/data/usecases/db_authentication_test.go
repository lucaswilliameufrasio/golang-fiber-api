// +test

package aucs_test

import (
	"errors"
	aucs "lucaswilliameufrasio/golang-fiber-api/src/data/usecases"
	ucs "lucaswilliameufrasio/golang-fiber-api/src/domain/usecases"
	fakeproto "lucaswilliameufrasio/golang-fiber-api/test/data/mocks"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	Faker faker.Faker
)

type SutTypes struct {
	sut                           ucs.Authentication
	encrypterSpy                  *fakeproto.FakeEncrypter
	loadUserByEmailRepositoryStub *fakeproto.FakeLoadUserByEmailRepository
	hashComparer                  *fakeproto.FakeHashComparer
}

func SUT(t *testing.T) SutTypes {
	Faker = faker.New()
	fakeEncrypter := &fakeproto.FakeEncrypter{}
	fakeLoadUserByEmailRepository := &fakeproto.FakeLoadUserByEmailRepository{}
	fakeHashComparer := &fakeproto.FakeHashComparer{}

	sut := aucs.NewDbAuthentication(fakeEncrypter, fakeLoadUserByEmailRepository, fakeHashComparer)

	return SutTypes{
		sut:                           sut,
		encrypterSpy:                  fakeEncrypter,
		loadUserByEmailRepositoryStub: fakeLoadUserByEmailRepository,
		hashComparer:                  fakeHashComparer,
	}
}

func TestCallLoadUserByEmailRepositoryCorrectly(t *testing.T) {
	instances := SUT(t)

	authenticationParams := ucs.AuthenticationParams{
		Email:    Faker.Internet().Email(),
		Password: Faker.Internet().Password(),
	}

	instances.loadUserByEmailRepositoryStub.LoadByEmail(authenticationParams.Email)

	_, err := instances.sut.Auth(authenticationParams)

	assert.Equal(t, err, nil)
}

func TestThrowIfLoadUserByEmailRepositoryThrows(t *testing.T) {
	instances := SUT(t)

	authenticationParams := ucs.AuthenticationParams{
		Email:    Faker.Internet().Email(),
		Password: Faker.Internet().Password(),
	}

	instances.loadUserByEmailRepositoryStub.LoadByEmailReturns(nil, errors.New("Generic"))

	result, err := instances.sut.Auth(authenticationParams)

	assert.Nil(t, result)
	assert.Error(t, err)
}
