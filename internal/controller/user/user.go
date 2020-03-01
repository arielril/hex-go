package user

import (
	"github.com/arielril/hexgo/internal/container"
	"github.com/arielril/hexgo/internal/container/user"
)

type UserController interface {
	CreateUser(user *user.User) error
	CreateUserAndSendEmail(user *user.User) error
	GetUserByEmail(email string) ([]*user.User, error)
	GetUserById(id string) (*user.User, error)
}

type userController struct {
	c *container.Container
}

func NewUserController(c *container.Container) UserController {
	return &userController{c}
}

func (ctrl *userController) CreateUser(user *user.User) error {
	return ctrl.c.UserService.CreateUser(user)
}

func (ctrl *userController) CreateUserAndSendEmail(user *user.User) error {
	return nil
}

func (ctrl *userController) GetUserByEmail(email string) ([]*user.User, error) {
	return ctrl.c.UserService.FindAllUsers()
}

func (ctrl *userController) GetUserById(id string) (*user.User, error) {
	return ctrl.c.UserService.FindUserById(id)
}
