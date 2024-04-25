package users

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateMock(t *testing.T) {
	repository := NewMockRepository(t)
	service := NewService(repository)

	updatedUser := User{
		ID:       1,
		Name:     "Updated User",
		Lastname: "Updated Lastname",
		Email:    "updated@example.com",
		Age:      30,
		Height:   1.80,
		Active:   true,
	}

	repository.On("Update", updatedUser.ID, updatedUser.Name, updatedUser.Lastname, updatedUser.Email, updatedUser.Age, updatedUser.Height, updatedUser.Active).Return(updatedUser, nil).Once()

	result, err := service.Update(updatedUser.ID, updatedUser.Name, updatedUser.Lastname, updatedUser.Email, updatedUser.Age, updatedUser.Height, updatedUser.Active)

	repository.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, updatedUser, result)
}

func TestDeleteMock(t *testing.T) {
	repository := NewMockRepository(t)
	service := NewService(repository)

	repository.On("Delete", uint(1)).Return(nil).Once()

	err := service.Delete(uint(1))

	repository.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestDeleteMockError(t *testing.T) {
	repository := NewMockRepository(t)
	service := NewService(repository)

	repository.On("Delete", uint(1)).Return(errors.New("unable to delete")).Once()

	err := service.Delete(uint(1))

	repository.AssertExpectations(t)
	assert.Error(t, err)
}

func TestStoreMock(t *testing.T) {
	repository := NewMockRepository(t)
	service := NewService(repository)

	storeUser := User{
		ID:        uint(2),
		Name:      "Store User",
		Lastname:  "Store Lastname",
		Email:     "stored@example.com",
		Age:       30,
		Height:    1.80,
		Active:    true,
		CreatedAt: "2021-01-01 00:00:00",
	}

	repository.On("LastId").Return(uint(1), nil).Once()
	repository.On("Store", storeUser.ID, storeUser.Name, storeUser.Lastname, storeUser.Email, mock.AnythingOfType("string"), storeUser.Age, storeUser.Height, storeUser.Active).Return(storeUser, nil).Once()

	result, err := service.Store(storeUser.Name, storeUser.Lastname, storeUser.Email, storeUser.Age, storeUser.Height, storeUser.Active)

	repository.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, storeUser, result)
}

func TestStoreMockWithLastIdError(t *testing.T) {
	repository := NewMockRepository(t)
	service := NewService(repository)

	storeUser := User{
		ID:        uint(2),
		Name:      "Store User",
		Lastname:  "Store Lastname",
		Email:     "stored@example.com",
		Age:       30,
		Height:    1.80,
		Active:    true,
		CreatedAt: "2021-01-01 00:00:00",
	}

	repository.On("LastId").Return(uint(0), errors.New("error getting last id")).Once()

	_, err := service.Store(storeUser.Name, storeUser.Lastname, storeUser.Email, storeUser.Age, storeUser.Height, storeUser.Active)

	repository.AssertExpectations(t)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "error getting last id")
}

func TestStoreMockWithError(t *testing.T) {
	repository := NewMockRepository(t)
	service := NewService(repository)

	storeUser := User{
		ID:        uint(2),
		Name:      "Store User",
		Lastname:  "Store Lastname",
		Email:     "stored@example.com",
		Age:       30,
		Height:    1.80,
		Active:    true,
		CreatedAt: "2021-01-01 00:00:00",
	}

	repository.On("LastId").Return(uint(1), nil).Once()
	repository.On("Store", storeUser.ID, storeUser.Name, storeUser.Lastname, storeUser.Email, mock.AnythingOfType("string"), storeUser.Age, storeUser.Height, storeUser.Active).Return(User{}, errors.New("unable to create")).Once()

	_, err := service.Store(storeUser.Name, storeUser.Lastname, storeUser.Email, storeUser.Age, storeUser.Height, storeUser.Active)

	repository.AssertExpectations(t)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "unable to create")
}

func TestGetAllMock(t *testing.T) {
	repository := NewMockRepository(t)
	service := NewService(repository)

	storedUsers := []User{
		{
			ID:        uint(1),
			Name:      "Store User 1",
			Lastname:  "Store Lastname 1",
			Email:     "stored1@example.com",
			Age:       31,
			Height:    1.81,
			Active:    true,
			CreatedAt: "2021-01-01 00:00:00",
		},
		{
			ID:        uint(2),
			Name:      "Store User 2",
			Lastname:  "Store Lastname 2",
			Email:     "stored2@example.com",
			Age:       32,
			Height:    1.82,
			Active:    true,
			CreatedAt: "2021-01-01 00:00:00",
		},
	}

	repository.On("GetAll").Return(storedUsers, nil).Once()

	users, err := service.GetAll()

	repository.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, users[0].ID, uint(1))
	assert.Equal(t, users[1].ID, uint(2))
}

func TestGetAllMockError(t *testing.T) {
	repository := NewMockRepository(t)
	service := NewService(repository)

	repository.On("GetAll").Return([]User{}, errors.New("error")).Once()

	_, err := service.GetAll()

	repository.AssertExpectations(t)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "error")
}

func TestGetByIdMock(t *testing.T) {
	repository := NewMockRepository(t)
	service := NewService(repository)

	storedUser := User{
		ID:        uint(1),
		Name:      "Store User 1",
		Lastname:  "Store Lastname 1",
		Email:     "stored1@example.com",
		Age:       31,
		Height:    1.81,
		Active:    true,
		CreatedAt: "2021-01-01 00:00:00",
	}

	repository.On("GetById", uint(1)).Return(storedUser, nil).Once()

	user, err := service.GetById(uint(1))

	repository.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, uint(1))
	assert.Equal(t, user.Name, "Store User 1")
}

func TestGetByIdMockError(t *testing.T) {
	repository := NewMockRepository(t)
	service := NewService(repository)

	repository.On("GetById", uint(1)).Return(User{}, errors.New("error")).Once()

	_, err := service.GetById(uint(1))

	repository.AssertExpectations(t)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "error")
}

func TestPatchMock(t *testing.T) {
	repository := NewMockRepository(t)
	service := NewService(repository)

	repository.On("Patch", uint(1), "test", 20).Return(User{}, nil).Once()

	us, err := service.Patch(uint(1), "test", 20)

	repository.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, us, User{})
}
