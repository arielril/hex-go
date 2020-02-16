package user

import (
	"github.com/arielril/hexgo/internal/container"
	"github.com/arielril/hexgo/internal/container/user"
)

type UserController interface {
	Create(user *user.User) error
	Get() ([]*user.User, error)
	GetById(id string) (*user.User, error)
}

type userController struct {
	c *container.Container
}

func NewUserController(c *container.Container) UserController {
	return &userController{c}
}

func (ctrl *userController) Create(user *user.User) error {
	return ctrl.c.UserService.CreateUser(user)
}

func (ctrl *userController) Get() ([]*user.User, error) {
	return ctrl.c.UserService.FindAllUsers()
}

func (ctrl *userController) GetById(id string) (*user.User, error) {
	return ctrl.c.UserService.FindUserById(id)
}
