package user

import (
	"time"

	// Model package
	. "github.com/arielril/hexgo/internal/container/model"

	"github.com/google/uuid"
)

// UserService interface
type UserService interface {
	CreateUser(user *User) error
	FindUserById(id string) (*User, error)
	FindUserByParams(params map[string]interface{}) ([]*User, error)
	FindAllUsers() ([]*User, error)
}

type userService struct {
	repo UserRepository
}

// NewUserService create a new user service
func NewUserService(repo UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

// CreateUser create a new user in the repository
func (us *userService) CreateUser(user *User) error {
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return us.repo.CreateUser(user)
}

// FindUserById search a user by it's ID
func (us *userService) FindUserById(id string) (*User, error) {
	params := make(map[string]interface{}, 0)
	params["ID"] = id
	user, err := us.repo.FindUser(params)

	return user[0], err
}

// FindUserByParams search the users using some params
func (us *userService) FindUserByParams(params map[string]interface{}) ([]*User, error) {
	return us.repo.FindUser(params)
}

// FindAllUsers gets all users from the repository
func (us *userService) FindAllUsers() ([]*User, error) {
	return us.repo.FindAllUsers()
}
