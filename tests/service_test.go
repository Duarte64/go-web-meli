package tests

import (
	"testing"

	"github.com/Duarte64/go-web-meli/internal/users"
	"github.com/Duarte64/go-web-meli/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateWithMock(t *testing.T) {
	mockUser := users.User{
		ID:        1,
		Name:      "Teste",
		Lastname:  "Update",
		Email:     "test@test.com",
		Age:       22,
		Height:    1.7,
		Active:    true,
		CreatedAt: "2024",
	}
	service, repository := createService(t)

	repository.Mock.On("Update",
		mock.AnythingOfType("uint"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("int"),
		mock.AnythingOfType("float64"),
		mock.AnythingOfType("bool"),
	).Return(mockUser, nil)

	user, err := service.Update(1, "test", "update", "test@test.com", 2, 1.7, true)

	assert.Nil(t, err)
	assert.Equal(t, int(user.ID), 1)
	assert.Equal(t, user.Name, "Teste")
}

func TestDeleteWithMock(t *testing.T) {
	service, repository := createService(t)

	repository.Mock.On("Delete",
		mock.AnythingOfType("uint"),
	).Return(nil)

	err := service.Delete(1)

	assert.Nil(t, err)
}

func createService(t *testing.T) (users.Service, *mocks.UserRepositoryMock) {
	t.Helper()
	repoMock := new(mocks.UserRepositoryMock)
	service := users.NewService(repoMock)
	return service, repoMock
}
