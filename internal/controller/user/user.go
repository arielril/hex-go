package user

import (
	"github.com/arielril/hexgo/internal/container"
	"github.com/arielril/hexgo/internal/container/model"
)

type UserController interface {
	CreateUser(user *model.User) error
	CreateUserAndSendEmail(user *model.User) error
	GetUserByEmail(email string) ([]*model.User, error)
	GetUserById(id string) (*model.User, error)
}

type userController struct {
	c *container.Container
}

func NewUserController(c *container.Container) UserController {
	return &userController{c}
}

func (ctrl *userController) CreateUser(user *model.User) error {
	return ctrl.c.UserService.CreateUser(user)
}

func (ctrl *userController) CreateUserAndSendEmail(user *model.User) error {
	return nil
}

func (ctrl *userController) GetUserByEmail(email string) ([]*model.User, error) {
	return ctrl.c.UserService.FindAllUsers()
}

func (ctrl *userController) GetUserById(id string) (*model.User, error) {
	return ctrl.c.UserService.FindUserById(id)
}
