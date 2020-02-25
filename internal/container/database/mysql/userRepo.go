package mysql

import (
	"github.com/arielril/hexgo/internal/container/user"
)

const table = "users"

type userRepository struct {
	*Persistence
}

func NewUserRepo() (user.UserRepository, error) {
	db, err := NewConnection()

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&user.User{})

	return &userRepository{db}, nil
}

func (r *userRepository) CreateUser(user *user.User) error {
	r.Table(table).Create(user)
	return nil
}

func (r *userRepository) FindUser(params map[string]interface{}) ([]*user.User, error) {
	results := r.Table(table).Find(make([]*user.User, 0), params).Value

	return results.([]*user.User), nil
}

func (r *userRepository) FindAllUsers() ([]*user.User, error) {
	results := r.Table(table).Find(make([]*user.User, 0)).Value
	return results.([]*user.User), nil
}
