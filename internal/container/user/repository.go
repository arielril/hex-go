package user

import (
	. "github.com/arielril/hexgo/internal/container/model"
)

type UserRepository interface {
	CreateUser(user *User) error
	FindUser(params map[string]interface{}) ([]*User, error)
	FindAllUsers() ([]*User, error)
}
