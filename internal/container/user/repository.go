package user

import (
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
