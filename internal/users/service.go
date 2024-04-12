package users

import "time"

type Service interface {
	GetAll() ([]User, error)
	GetById(id uint) (User, error)
	Store(name, lastname, email string, age int, height float64, active bool) (User, error)
	Update(id uint, name, lastname, email string, age int, height float64, active bool) (User, error)
	Patch(id uint, lastname string, age int) (User, error)
	Delete(id uint) error
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]User, error) {
	us, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return us, nil
}

func (s *service) GetById(id uint) (User, error) {
	us, err := s.repository.GetById(id)
	if err != nil {
		return User{}, err
	}

	return us, nil
}

func (s *service) Store(name, lastname, email string, age int, height float64, active bool) (User, error) {
	lastId, err := s.repository.LastId()
	date := time.Now().String()
	if err != nil {
		return User{}, err
	}

	lastId++

	u, err := s.repository.Store(lastId, name, lastname, email, date, age, height, active)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func (s *service) Update(id uint, name, lastname, email string, age int, height float64, active bool) (User, error) {
	return s.repository.Update(id, name, lastname, email, age, height, active)
}

func (s *service) Patch(id uint, lastname string, age int) (User, error) {
	return s.repository.Patch(id, lastname, age)
}

func (s *service) Delete(id uint) error {
	if err := s.repository.Delete(id); err != nil {
		return err
	}

	return nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
