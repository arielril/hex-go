package user

import (
	// Model package
	. "github.com/arielril/hexgo/internal/container/model"
)

// UserRepository interface
type UserRepository interface {
	CreateUser(user *User) error
	FindUser(params map[string]interface{}) ([]*User, error)
	FindAllUsers() ([]*User, error)
}
