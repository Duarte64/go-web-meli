package users

var us []User
var lastId uint

type repository struct{}

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
	LastId() (*uint, error)
}

func (r *repository) GetAll() ([]User, error) {
	return us, nil
}

func (r *repository) GetById(id uint) (User, error) {
	for _, user := range us {
		if user.ID == id {
			return user, nil
		}
	}

	return User{}, &NotFoundError{}
}

func (r *repository) Store(id uint, name, lastname, email, createdAt string, age int, height float64, active bool) (User, error) {
	user := User{id, name, lastname, email, age, height, active, createdAt}
	us = append(us, user)
	return user, nil
}

func (r *repository) Update(id uint, name, lastname, email string, age int, height float64, active bool) (User, error) {
	updatedUser := User{Name: name, Lastname: lastname, Email: email, Age: age, Height: height, Active: active}
	for index, user := range us {
		if user.ID == id {
			updatedUser.ID = user.ID
			updatedUser.CreatedAt = user.CreatedAt
			us[index] = updatedUser
			return updatedUser, nil
		}
	}
	return User{}, &NotFoundError{}
}

func (r *repository) Patch(id uint, lastname string, age int) (User, error) {
	for index, user := range us {
		if user.ID == id {
			if lastname != "" {
				user.Lastname = lastname
			}
			if age != 0 {
				user.Age = age
			}
			us[index] = user
			return user, nil
		}
	}
	return User{}, &NotFoundError{}
}

func (r *repository) LastId() (*uint, error) {
	return &lastId, nil
}

func (r *repository) Delete(id uint) error {
	for index, user := range us {
		if user.ID == id {
			us = append(us[:index], us[index+1:]...)
			return nil
		}
	}
	return &NotFoundError{}
}

func NewRepository() Repository {
	return &repository{}
}
