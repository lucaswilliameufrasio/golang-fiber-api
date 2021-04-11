// +test

package aucs_test

import (
	"errors"
	fakeproto "lucaswilliameufrasio/golang-fiber-api/src/data/protocols/protocolsfakes"
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
	encrypterSpy                  *fakeproto.FakeEncrypter
	loadUserByEmailRepositoryStub *fakeproto.FakeLoadUserByEmailRepository
	hashComparer                  *fakeproto.FakeHashComparer
}

func SUT(t *testing.T) SutTypes {
	Faker = faker.New()
	MockController = gomock.NewController(t)
	defer MockController.Finish()
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
