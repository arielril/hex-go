package mysql

import (
	"github.com/arielril/hexgo/internal/container/model"
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

	db.AutoMigrate(&model.User{})

	return &userRepository{db}, nil
}

func (r *userRepository) CreateUser(user *model.User) error {
	r.Table(table).Create(user)
	return nil
}

func (r *userRepository) FindUser(params map[string]interface{}) ([]*model.User, error) {
	results := r.Table(table).Find(make([]*model.User, 0), params).Value

	return results.([]*model.User), nil
}

func (r *userRepository) FindAllUsers() ([]*model.User, error) {
	results := r.Table(table).Find(make([]*model.User, 0)).Value
	return results.([]*model.User), nil
}
