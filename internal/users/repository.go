package users

import (
	"github.com/Duarte64/go-web-meli/pkg/store"
)

type repository struct {
	db store.Store
}

type NotFoundError struct{}

func (n *NotFoundError) Error() string {
	return "Usuário não encontrado"
}

type Repository interface {
	GetAll() ([]User, error)
	GetById(id uint) (User, error)
	Delete(id uint) error
	Store(id uint, name, lastname, email, createdAt string, age int, height float64, active bool) (User, error)
	Update(id uint, name, lastname, email string, age int, height float64, active bool) (User, error)
	Patch(id uint, lastname string, age int) (User, error)
	LastId() (uint, error)
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) LastId() (uint, error) {
	var us []User
	if err := r.db.Read(&us); err != nil {
		return 0, err
	}
	if len(us) == 0 {
		return 0, nil
	}

	maiorId := us[0].ID
	for _, user := range us {
		if user.ID > maiorId {
			maiorId = user.ID
		}
	}
	return maiorId, nil
}

func (r *repository) Store(id uint, name, lastname, email, createdAt string, age int, height float64, active bool) (User, error) {
	var us []User
	if err := r.db.Read(&us); err != nil {
		return User{}, err
	}
	u := User{id, name, lastname, email, age, height, active, createdAt}
	us = append(us, u)
	if err := r.db.Write(us); err != nil {
		return User{}, err
	}
	return u, nil
}

func (r *repository) Update(id uint, name, lastname, email string, age int, height float64, active bool) (User, error) {
	var us []User
	if err := r.db.Read(&us); err != nil {
		return User{}, err
	}
	updatedUser := User{Name: name, Lastname: lastname, Email: email, Age: age, Height: height, Active: active}
	for index, user := range us {
		if user.ID == id {
			updatedUser.ID = user.ID
			updatedUser.CreatedAt = user.CreatedAt
			us[index] = updatedUser
			if err := r.db.Write(us); err != nil {
				return User{}, err
			} else {
				return updatedUser, nil
			}
		}
	}
	return User{}, &NotFoundError{}
}

func (r *repository) Delete(id uint) error {
	var us []User
	if err := r.db.Read(&us); err != nil {
		return err
	}
	for index, user := range us {
		if user.ID == id {
			us = append(us[:index], us[index+1:]...)
			if err := r.db.Write(us); err != nil {
				return err
			} else {
				return nil
			}
		}
	}
	return &NotFoundError{}
}

func (r *repository) GetAll() ([]User, error) {
	var us []User
	if err := r.db.Read(&us); err != nil {
		return []User{}, err
	}
	return us, nil
}

func (r *repository) GetById(id uint) (User, error) {
	var us []User
	if err := r.db.Read(&us); err != nil {
		return User{}, err
	}
	for _, user := range us {
		if user.ID == id {
			return user, nil
		}
	}

	return User{}, &NotFoundError{}
}

func (r *repository) Patch(id uint, lastname string, age int) (User, error) {
	var us []User
	if err := r.db.Read(&us); err != nil {
		return User{}, err
	}
	for index, user := range us {
		if user.ID == id {
			if lastname != "" {
				user.Lastname = lastname
			}
			if age != 0 {
				user.Age = age
			}
			us[index] = user
			if err := r.db.Write(us); err != nil {
				return User{}, err
			} else {
				return us[index], nil
			}
		}
	}
	return User{}, &NotFoundError{}
}
