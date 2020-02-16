package user

import (
	"fmt"
	"time"
)

type User struct {
	ID           string
	FullName     string
	EmailAddress string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserRepository interface {
	CreateUser(user *User) error
	FindUser(params map[string]interface{}) ([]*User, error)
	FindAllUsers() ([]*User, error)
}

type userRepository struct {
	database interface{}
}

func NewUserRepo(db interface{}) UserRepository {
	return &userRepository{
		database: db,
	}
}

func (r *userRepository) CreateUser(user *User) error {
	fmt.Printf("Inserting user %#v\n", user)
	return nil
}

func (r *userRepository) FindUser(params map[string]interface{}) ([]*User, error) {
	return nil, nil
}

func (r *userRepository) FindAllUsers() ([]*User, error) {
	return nil, nil
}
