package user

import (
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(user *User) error
	FindUserById(id string) (*User, error)
	FindUserByParams(params map[string]interface{}) ([]*User, error)
	FindAllUsers() ([]*User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (us *userService) CreateUser(user *User) error {
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return us.repo.CreateUser(user)
}

func (us *userService) FindUserById(id string) (*User, error) {
	params := make(map[string]interface{}, 0)
	params["ID"] = id
	user, err := us.repo.FindUser(params)

	return user[0], err
}

func (us *userService) FindUserByParams(params map[string]interface{}) ([]*User, error) {
	return us.repo.FindUser(params)
}

func (us *userService) FindAllUsers() ([]*User, error) {
	return us.repo.FindAllUsers()
}
