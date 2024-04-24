package mocks

import (
	"github.com/Duarte64/go-web-meli/internal/users"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func NewUserRepositoryMock() users.Repository {
	return &UserRepositoryMock{}
}

func (p *UserRepositoryMock) GetById(id uint) (users.User, error) {
	args := p.Called(id)
	// argumentos: [[]entities.Product{}, error]
	return args.Get(0).(users.User), args.Error(1)
}

func (p *UserRepositoryMock) GetAll() ([]users.User, error) {
	return []users.User{}, nil
}

func (p *UserRepositoryMock) Delete(id uint) error {
	args := p.Called(id)

	return args.Error(0)
}

func (p *UserRepositoryMock) Store(id uint, name, lastname, email, createdAt string, age int, height float64, active bool) (users.User, error) {
	return users.User{}, nil
}

func (p *UserRepositoryMock) Update(id uint, name, lastname, email string, age int, height float64, active bool) (users.User, error) {
	args := p.Called(id, name, lastname, email, age, height, active)

	return args.Get(0).(users.User), args.Error(1)
}

func (p *UserRepositoryMock) Patch(id uint, lastname string, age int) (users.User, error) {
	// args := p.Called(id, lastname, age)

	return users.User{}, nil
}

func (p *UserRepositoryMock) LastId() (uint, error) {
	return 1, nil
}
